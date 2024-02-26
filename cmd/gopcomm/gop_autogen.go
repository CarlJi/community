package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/goplus/yap"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	_ "github.com/joho/godotenv/autoload"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
)

const (
	layoutUS   = "January 2, 2006"
	limitConst = 10
)

type community_yap struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:28
func (this *community_yap) MainEntry() {
//line cmd/gopcomm/community_yap.gox:28:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:29:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:30:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:31:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:35:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:37:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:38:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:41:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:42:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:45:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:46:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:49:1
	this.Get("/demo", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:50:1
		ctx.Yap__1("demo", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:53:1
	this.Get("/signin", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:54:1
		ctx.Yap__1("signin", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:57:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:60:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:61:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:62:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:63:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:64:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:65:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:66:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:68:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:72:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:73:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:74:1
		articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:75:1
		ctx.Yap__1("article", map[string]interface {
		}{"UserId": userId, "User": user, "Article": string(articleJson)})
	})
//line cmd/gopcomm/community_yap.gox:82:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:83:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:84:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:85:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:91:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:92:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:94:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:95:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:96:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:99:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:100:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:101:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:102:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:103:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:104:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:105:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:107:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:111:1
		items, next, _ := this.community.GetArticlesByUid(todo, id, core.MarkBegin, limitConst)
//line cmd/gopcomm/community_yap.gox:112:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:113:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:114:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": string(userClaimJson), "User": user, "Items": string(itemsJson), "UserId": userId, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:123:1
	this.Get("/userUnlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:124:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:125:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:126:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:127:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusTemporaryRedirect)
//line cmd/gopcomm/community_yap.gox:128:1
			return
		}
//line cmd/gopcomm/community_yap.gox:130:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:131:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:132:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:133:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:134:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:135:1
		default:
//line cmd/gopcomm/community_yap.gox:136:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:138:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:139:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:140:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:142:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/userEdit", http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:144:1
	this.Get("/userEdit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:145:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:146:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:147:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:149:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:150:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:151:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:153:1
		fullUser, _ := gac.GetUser()
//line cmd/gopcomm/community_yap.gox:154:1
		appInfo, _ := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:155:1
		appInfoStr, _ := json.Marshal(*appInfo)
//line cmd/gopcomm/community_yap.gox:156:1
		binds, _ := json.Marshal(gac.GetProviderBindStatus())
//line cmd/gopcomm/community_yap.gox:157:1
		currentUser, _ := json.Marshal(fullUser)
//line cmd/gopcomm/community_yap.gox:158:1
		user := gac.GetUserSimple()
//line cmd/gopcomm/community_yap.gox:159:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"UserId": user.Id, "User": user, "CurrentUser": string(currentUser), "Application": string(appInfoStr), "Binds": string(binds)})
	})
//line cmd/gopcomm/community_yap.gox:168:1
	this.Post("/update/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:169:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:170:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:171:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("name")}
//line cmd/gopcomm/community_yap.gox:182:1
		_, err = this.community.UpdateUserById(fmt.Sprintf("%s/%s", user.Owner, user.Name), user)
//line cmd/gopcomm/community_yap.gox:183:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:184:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:185:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:190:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:195:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:196:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:197:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:198:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:199:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:200:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:201:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:202:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:204:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:207:1
		ctx.Yap__1("edit", map[string]interface {
		}{"User": user, "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:213:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:214:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:215:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:216:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:217:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:218:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:219:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:224:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:225:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:226:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:231:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:238:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:239:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:240:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:241:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:242:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:243:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:244:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:245:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:247:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:248:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:249:1
			pageInt = 1
		}
//line cmd/gopcomm/community_yap.gox:251:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:252:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:253:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "count": count, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:259:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "count": count, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:267:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:268:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:269:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:270:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:271:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:272:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:273:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:278:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:279:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:280:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:285:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})

//line cmd/gopcomm/community_yap.gox:292:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:293:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:295:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:296:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:297:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:298:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:299:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:300:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:302:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:306:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:307:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:308:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:316:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:317:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:318:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:319:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:321:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:322:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:323:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:326:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:327:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:335:1
	this.Get("/userArticles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:336:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:337:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:338:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:340:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:341:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:342:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:344:1
		items, next, _ := this.community.GetArticlesByUid(todo, id, from, limitInt)
//line cmd/gopcomm/community_yap.gox:345:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "next": next})
	})
//line cmd/gopcomm/community_yap.gox:352:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:353:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:361:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:363:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:364:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:365:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:366:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:367:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:368:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:370:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:374:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:375:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:376:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:385:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:386:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:387:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:388:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:389:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:390:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:391:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:392:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:394:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:398:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:399:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:400:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:401:1
			if
//line cmd/gopcomm/community_yap.gox:401:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:402:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:403:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:405:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:406:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:407:1
			ctx.Yap__1("edit", map[string]interface {
			}{"UserId": userId, "User": user, "Article": string(articleJson)})
		}
	})
//line cmd/gopcomm/community_yap.gox:415:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:416:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:417:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:418:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:419:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:424:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:431:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:432:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:433:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:434:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:435:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:436:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:437:1
		trans, _ := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:439:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:440:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:441:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:442:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:447:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:448:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:449:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:450:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:455:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract}, Content: content, Trans: trans}
//line cmd/gopcomm/community_yap.gox:469:1
		if trans {
//line cmd/gopcomm/community_yap.gox:470:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:473:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:474:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:475:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:480:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:486:1
	this.Get("/tranArticle", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:487:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:489:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:490:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:491:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:495:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:504:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:505:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:506:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:507:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:508:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:513:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:520:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:521:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:523:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:525:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:528:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:529:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:530:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:531:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:532:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:533:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:538:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:544:1
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:545:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:546:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:547:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:548:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:549:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:550:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:555:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:556:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:557:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:558:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:559:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:564:1
				return
			}
//line cmd/gopcomm/community_yap.gox:566:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:568:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:569:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:570:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:575:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:576:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:582:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:583:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:586:1
	this.Post("/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:587:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:588:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:590:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:591:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:597:1
		if
//line cmd/gopcomm/community_yap.gox:597:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:598:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:604:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:610:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:615:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:616:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:617:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:618:1
			return
		}
//line cmd/gopcomm/community_yap.gox:621:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:622:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:623:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:626:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:628:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:629:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:632:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:633:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:634:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:635:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:638:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:639:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:640:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:641:1
			return
		}
//line cmd/gopcomm/community_yap.gox:644:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:645:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:646:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:649:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:652:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:653:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:654:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:655:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:657:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:658:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:659:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:660:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:661:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:664:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:667:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:668:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:669:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:672:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:673:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:676:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:679:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:681:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:682:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:683:1
				if
//line cmd/gopcomm/community_yap.gox:683:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:684:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:685:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:689:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community_yap))
}
