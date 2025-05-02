package modals

import "sync"

type PageInfoModal struct {
	Url               string
	HeadingProperties []Property
	HtmlVersion       string `json:"htmlVersion"`
	Title             string `json:"title"`
	//ServiceTime    int64     `json:"serviceTime"`
	//WebExtractTime int64     `json:"webExtractTime"`
	NoOfInternalLinks     int  `json:"noOfInternalLinks"`
	NoOfExternalLinks     int  `json:"noOfExternalLinks"`
	NoOfInaccessibleLinks int  `json:"noOfInaccessibleLinks"`
	HasLogin              bool `json:"hasLogin"`
}

type Property struct {
	propertyName        string
	numberOfOccurrences int
}

type PageInfoModalManager struct {
	pageInfoModal PageInfoModal
	lock          sync.Mutex
}

func NewPageInfoModalManager() PageInfoModalManager {

	properties := make([]Property, 0)
	return PageInfoModalManager{
		pageInfoModal: PageInfoModal{
			HeadingProperties: properties,
		},
	}
}

func (modalManager *PageInfoModalManager) SetHtmlVersion(htmlVersion string) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.HtmlVersion = htmlVersion
}

func (modalManager *PageInfoModalManager) SetTitle(title string) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.Title = title
}

func (modalManager *PageInfoModalManager) SetUrl(url string) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.Url = url
}

func (modalManager *PageInfoModalManager) SetHasLogin(hasLogin bool) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.HasLogin = hasLogin
}
