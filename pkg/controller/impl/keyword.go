package impl

import (
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"github.com/iohub/ahocorasick"
	"log"
	"sync"
	"time"
)

type KeywordControllerImpl struct {
	matcherMap sync.Map
	key string
}

var _ controller.KeywordController = (*KeywordControllerImpl)(nil)

func NewKeywordController() *KeywordControllerImpl {
	key := "keyword"
	matcher := buildMatcher()
	matcherMap := sync.Map{}
	matcherMap.Store(key, matcher)
	keywordController := &KeywordControllerImpl{matcherMap:matcherMap}
	keywordController.syncMatcher()
	return keywordController
}

func (p *KeywordControllerImpl) Search(word string) bool {
	seq := []byte(word)
	if value, ok := p.matcherMap.Load(p.key); ok {
		matcher := value.(*cedar.Matcher)
		resp := matcher.Match(seq)
		if resp.HasNext() {
			return true
		}
	}
	return false
}

func buildMatcher() *cedar.Matcher {
	matcher := cedar.NewMatcher()
	words := impl.DefaultKeywordDAO.GetKeywords()
	for _, word := range words {
		matcher.Insert([]byte(word.Word), word.Id)
	}
	matcher.Compile()
	return matcher
}

func (p *KeywordControllerImpl) syncMatcher() {
	go func() {
		for {
			log.Println("update matcher")
			matcher := buildMatcher()
			p.matcherMap.Store(p.key, matcher)
			time.Sleep(time.Duration(configs.DefaultConfig.Keyword.Tick) * time.Second)
		}
	}()
}
