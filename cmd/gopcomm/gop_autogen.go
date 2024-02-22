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
	labelConst = "article"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:29
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:29:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:30:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:31:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:32:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:36:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:38:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:39:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:42:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:43:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:44:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:45:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:46:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:47:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:48:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:49:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:51:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:54:1
		ctx.Yap__1("4xx", map[string]interface {
		}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:60:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:61:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:62:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:63:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:64:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:65:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:66:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:67:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:69:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:72:1
		ctx.Yap__1("5xx", map[string]interface {
		}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:79:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:81:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:82:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:83:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:84:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:85:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:86:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:87:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:89:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:93:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
//line cmd/gopcomm/community_yap.gox:94:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:95:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:97:1
		articlesJson, err := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:98:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:99:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:101:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:112:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:115:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:116:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:117:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:118:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:119:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:120:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:121:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:123:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:127:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:128:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:129:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:130:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:132:1
		articleJson, err := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:133:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:134:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:136:1
		ctx.Yap__1("article", map[string]interface {
		}{"UserId": userId, "User": user, "Article": string(articleJson)})
	})
//line cmd/gopcomm/community_yap.gox:143:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:144:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:145:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:146:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:147:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:148:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:149:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:150:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:152:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:155:1
		ctx.Yap__1("edit", map[string]interface {
		}{"User": user, "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:161:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:162:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:163:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:164:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:165:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:166:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:167:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:168:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:170:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:174:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:175:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:176:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:177:1
			editable, err := this.community.CanEditable(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:178:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:179:1
				xLog.Error("can editable error:", err)
//line cmd/gopcomm/community_yap.gox:180:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:182:1
			if !editable {
//line cmd/gopcomm/community_yap.gox:183:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:184:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:186:1
			article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:187:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:188:1
				xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:189:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:191:1
			articleJson, err := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:192:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:193:1
				xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:194:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:196:1
			ctx.Yap__1("edit", map[string]interface {
			}{"UserId": userId, "User": user, "Article": string(articleJson)})
		}
	})
//line cmd/gopcomm/community_yap.gox:204:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:205:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:206:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:207:1
		if label == "" {
//line cmd/gopcomm/community_yap.gox:208:1
			label = "article"
		}
//line cmd/gopcomm/community_yap.gox:212:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:213:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:214:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:215:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:216:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:217:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:218:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:220:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:224:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
//line cmd/gopcomm/community_yap.gox:225:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:226:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:228:1
		articlesJson, err := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:229:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:230:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:232:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Value": searchValue, "Next": next, "Tab": label})
	})
//line cmd/gopcomm/community_yap.gox:245:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:246:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:247:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:248:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:249:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:250:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:255:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:261:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:262:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:263:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:264:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:265:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:266:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:267:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:272:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:273:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:274:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:279:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:286:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:287:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:288:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:289:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:290:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:292:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:293:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:294:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:297:1
		articles, next, err := this.community.ListArticle(todo, from, limitInt, searchValue, label)
//line cmd/gopcomm/community_yap.gox:298:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:299:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:300:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:305:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:313:1
	this.Post("/api/article/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:314:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:315:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:316:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:317:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:318:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:319:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:320:1
		trans, err := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:321:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:322:1
			xLog.Error("parse bool error:", err)
		}
//line cmd/gopcomm/community_yap.gox:325:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:326:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:327:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:328:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:333:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:334:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:335:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:336:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:341:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract, Label: label}, Content: content, Trans: trans}
//line cmd/gopcomm/community_yap.gox:356:1
		if trans {
//line cmd/gopcomm/community_yap.gox:357:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:360:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:361:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:362:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:367:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:376:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:377:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:379:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:380:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:381:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:384:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:385:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:386:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:387:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:388:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:389:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:390:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:392:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:396:1
		items, next, err := this.community.GetArticlesByUid(todo, id, core.MarkBegin, limitConst)
//line cmd/gopcomm/community_yap.gox:397:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:398:1
			xLog.Error("get article list error:", err)
		}
//line cmd/gopcomm/community_yap.gox:400:1
		userClaimJson, err := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:401:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:402:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:404:1
		itemsJson, err := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:405:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:406:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:408:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": string(userClaimJson), "User": user, "Items": string(itemsJson), "UserId": userId, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:418:1
	this.Get("/user/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:419:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:420:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:421:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:423:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:424:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:425:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:427:1
		fullUser, err := gac.GetUser()
//line cmd/gopcomm/community_yap.gox:428:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:429:1
			xLog.Error("get user error:", err)
//line cmd/gopcomm/community_yap.gox:430:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:432:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:433:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:434:1
			xLog.Error("get application info error:", err)
//line cmd/gopcomm/community_yap.gox:435:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:437:1
		appInfoStr, err := json.Marshal(*appInfo)
//line cmd/gopcomm/community_yap.gox:438:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:439:1
			xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:440:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:442:1
		binds, err := json.Marshal(gac.GetProviderBindStatus())
//line cmd/gopcomm/community_yap.gox:443:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:444:1
			xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:445:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:447:1
		currentUser, err := json.Marshal(fullUser)
//line cmd/gopcomm/community_yap.gox:448:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:449:1
			xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:450:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:452:1
		user := gac.GetUserSimple()
//line cmd/gopcomm/community_yap.gox:453:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"UserId": user.Id, "User": user, "CurrentUser": string(currentUser), "Application": string(appInfoStr), "Binds": string(binds)})
	})
//line cmd/gopcomm/community_yap.gox:465:1
	this.Put("/api/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:466:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:467:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:468:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("name")}
//line cmd/gopcomm/community_yap.gox:479:1
		_, err = this.community.UpdateUserById(fmt.Sprintf("%s/%s", user.Owner, user.Name), user)
//line cmd/gopcomm/community_yap.gox:480:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:481:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:482:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:487:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:492:1
	this.Get("/api/user/unlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:493:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:494:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:495:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:496:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusTemporaryRedirect)
//line cmd/gopcomm/community_yap.gox:497:1
			return
		}
//line cmd/gopcomm/community_yap.gox:499:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:500:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:501:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:502:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:503:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:504:1
		default:
//line cmd/gopcomm/community_yap.gox:505:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:507:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:508:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:509:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:511:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/user/edit", http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:514:1
	this.Get("/api/user/:id/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:515:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:516:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:517:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:519:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:520:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:521:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:523:1
		items, next, err := this.community.GetArticlesByUid(todo, id, from, limitInt)
//line cmd/gopcomm/community_yap.gox:524:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:525:1
			xLog.Error("get article list error:", err)
//line cmd/gopcomm/community_yap.gox:526:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error(), "total": 0})
		}
//line cmd/gopcomm/community_yap.gox:532:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "next": next})
	})
//line cmd/gopcomm/community_yap.gox:539:1
	this.Get("/api/user/:id/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:540:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:541:1
		uid := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:542:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:543:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:545:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:546:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:547:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:549:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:550:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:551:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "total": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:557:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:558:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:559:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "total": total, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:565:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "total": total, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:576:1
	this.Delete("/api/media/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:577:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:578:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:579:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:580:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:581:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:582:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:587:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:588:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:589:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:594:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:601:1
	this.Get("/api/translation/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:602:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:603:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:604:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:605:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:609:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:617:1
	this.Post("/api/translation", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:618:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:619:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:620:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:621:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:626:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:632:1
	this.Get("/api/media/:id/url", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:633:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:634:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:635:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:636:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:637:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:642:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:648:1
	this.Get("/api/video/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:649:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:650:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:651:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:652:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:653:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:654:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:659:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:660:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:661:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:662:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:663:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:668:1
				return
			}
//line cmd/gopcomm/community_yap.gox:670:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:672:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:673:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:674:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:679:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:680:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:686:1
	this.Post("/api/media", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:687:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:690:1
	this.Post("/api/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:691:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:692:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:694:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:695:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:701:1
		if
//line cmd/gopcomm/community_yap.gox:701:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:702:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:708:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:717:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:722:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:723:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:724:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:725:1
			return
		}
//line cmd/gopcomm/community_yap.gox:728:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:729:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:730:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:733:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:735:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:736:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:739:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:740:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:741:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:742:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:745:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:746:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:747:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:748:1
			return
		}
//line cmd/gopcomm/community_yap.gox:751:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:752:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:753:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:756:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:759:1
	this.Get("/login/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:760:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:761:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:762:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:764:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:765:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:766:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:767:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:768:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:771:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:774:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:775:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:776:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:779:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:780:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:783:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:786:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:788:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:789:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:790:1
				if
//line cmd/gopcomm/community_yap.gox:790:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:791:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:792:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:796:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
