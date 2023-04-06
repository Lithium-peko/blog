package service

import (
	"blog/config"
	"blog/dao"
	"blog/model"
	"blog/model/dto"
	"blog/model/req"
	"blog/model/resp"
	"blog/utils"
	"blog/utils/r"
	"github.com/gin-contrib/sessions"
	"go.uber.org/zap"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct{}

// 登录
func (*User) Login(c *gin.Context, username, password string) (loginVo resp.LoginVO, code int) {
	// 检查用户是否存在
	userAuth := dao.GetOne(model.UserAuth{}, "username", username)
	if userAuth.ID == 0 {
		return loginVo, r.ERROR_USER_NOT_EXIST
	}
	// 检查密码是否正确
	if !utils.Encryptor.BcryptCheck(password, userAuth.Password) {
		return loginVo, r.ERROR_PASSWORD_WRONG
	}
	// 获取用户详细信息 DTO
	userDetailDTO := convertUserDetailDTO(userAuth, c)

	// 登录信息正确, 生成 Token
	// TODO: 目前只给用户设定一个角色, 获取第一个值就行, 后期优化: 给用户设置多个角色
	// UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	uuid := utils.Encryptor.MD5(userDetailDTO.IpAddress + userDetailDTO.Browser + userDetailDTO.OS)
	token, err := utils.GetJWT().GenToken(userAuth.ID, userDetailDTO.RoleLabels[0], uuid)
	if err != nil {
		utils.Logger.Info("登录时生成 Token 错误: ", zap.Error(err))
		return loginVo, r.ERROR_TOKEN_CREATE
	}
	userDetailDTO.Token = token
	// 更新用户验证信息: ip 信息 + 上次登录时间
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{ID: userAuth.ID},
		IpAddress:     userDetailDTO.IpAddress,
		IpSource:      userDetailDTO.IpSource,
		LastLoginTime: userDetailDTO.LastLoginTime,
	}, "ip_address", "ip_source", "last_login_time")

	// 保存用户信息到 Session 和 Redis 中
	session := sessions.Default(c)
	// ! session 中只能存储字符串
	sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{UserDetailDTO: userDetailDTO})
	session.Set(KEY_USER+uuid, sessionInfoStr) // ! 确实设置到 redis 中, 但是获取不到
	utils.Redis.Set(KEY_USER+uuid, sessionInfoStr, time.Duration(config.Cfg.Session.MaxAge)*time.Second)
	session.Save()

	return userDetailDTO.LoginVO, r.OK
}

// 退出登录
func (*User) LoginOut(c *gin.Context) {
	uuid := utils.GetFromContext[string](c, "uuid")
	session := sessions.Default(c)
	session.Delete(KEY_USER + uuid)
	session.Save()
	utils.Redis.Del(KEY_USER + uuid) // 删除 redis 中缓存
}

// 用户注册
func (*User) Register(req req.Register) (code int) {
	// 检查验证码是否正确
	if req.Code != utils.Redis.GetVal(KEY_CODE+req.Username) {
		return r.ERROR_VERIFICATION_CODE
	}

	// 检查用户名已存在, 则该账号已经注册过
	if exist := checkUserExistByName(req.Username); exist {
		return r.ERROR_USER_NAME_USED
	}

	userInfo := &model.UserInfo{
		Email:    req.Username,
		Nickname: "用户" + req.Username,
		// blogConfig 中配置的默认用户头像
		Avatar: blogInfoService.GetBlogConfig().UserAvatar,
	}
	dao.Create(&userInfo)
	// 设置用户默认角色
	dao.Create(&model.UserRole{
		UserId: userInfo.ID,
		RoleId: 3, // 默认角色是 "测试"
	})
	dao.Create(&model.UserAuth{
		UserInfoId:    userInfo.ID,
		Username:      req.Username,
		Password:      utils.Encryptor.BcryptHash(req.Password),
		LoginType:     1,
		LastLoginTime: time.Now(), // 注册时会更新 "上次登陆时间"
	})
}

// 检查用户是否存在
func checkUserExistByName(username string) bool {
	existUser := dao.GetOne(model.UserAuth{}, "username = ?", username)
	return existUser.ID != 0
}

// 转化 UserDetailDTO
func convertUserDetailDTO(userAuth model.UserAuth, c *gin.Context) dto.UserDetailDTO {
	// 获取 IP 相关信息 FIXME: 好像无法读取到 ip 信息
	ipAddress := utils.IP.GetIpAddress(c)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)
	browser, os := "unknown", "unknown"

	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}

	// 获取用户详细信息
	userInfo := dao.GetOne(&model.UserInfo{}, "id", userAuth.UserInfoId)
	// FIXME: 获取该用户对应的角色, 没有角色默认是 "test"
	roleLabels := roleDao.GetLabelsByUserInfoId(userInfo.ID)
	if len(roleLabels) == 0 {
		roleLabels = append(roleLabels, "test")
	}
	// 用户点赞 Set
	articleLikeSet := utils.Redis.SMembers(KEY_ARTICLE_USER_LIKE_SET + strconv.Itoa(userInfo.ID))
	commentLikeSet := utils.Redis.SMembers(KEY_COMMENT_USER_LIKE_SET + strconv.Itoa(userInfo.ID))

	return dto.UserDetailDTO{
		LoginVO: resp.LoginVO{
			ID:             userAuth.ID,
			UserInfoId:     userInfo.ID,
			Email:          userInfo.Email,
			LoginType:      userAuth.LoginType,
			Username:       userAuth.Username,
			Nickname:       userInfo.Nickname,
			Avatar:         userInfo.Avatar,
			Intro:          userInfo.Intro,
			Website:        userInfo.Website,
			IpAddress:      ipAddress,
			IpSource:       ipSource,
			LastLoginTime:  time.Now(),
			ArticleLikeSet: articleLikeSet,
			CommentLikeSet: commentLikeSet,
		},
		Password:   userAuth.Password,
		RoleLabels: roleLabels,
		IsDisable:  userInfo.IsDisable,
		Browser:    browser,
		OS:         os,
	}
}
