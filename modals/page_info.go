package modals

import "sync"

type PageInfoModal struct {
	Url                       string
	HeadingProperties         []Property
	HtmlVersion               string
	Title                     string
	NoOfInternalLinks         int
	NoOfExternalLinks         int
	NoOfInaccessibleLinks     int
	HasLogin                  bool
	InAccessibleLinksMetaInfo []string
}

type Property struct {
	PropertyName        string
	NumberOfOccurrences int
}

type PageInfoModalManager struct {
	pageInfoModal PageInfoModal
	lock          sync.Mutex
}

func NewPageInfoModalManager() *PageInfoModalManager {

	properties := make([]Property, 0)
	return &PageInfoModalManager{
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

func (modalManager *PageInfoModalManager) SetNoOfInternalLinks(noOfInternalLinks int) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.NoOfInternalLinks = noOfInternalLinks
}

func (modalManager *PageInfoModalManager) SetNoOfExternalLinks(noOfExternalLinks int) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.NoOfExternalLinks = noOfExternalLinks
}

func (modalManager *PageInfoModalManager) SetNoOfInaccessibleLinks(noOfInaccessibleLinks int) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.NoOfInaccessibleLinks = noOfInaccessibleLinks
}

func (modalManager *PageInfoModalManager) SetInAccessibleMetaInfoLinks(inAccessibleUrl string) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.InAccessibleLinksMetaInfo = append(
		modalManager.pageInfoModal.InAccessibleLinksMetaInfo,
		inAccessibleUrl,
	)
}

func (modalManager *PageInfoModalManager) SetHeadingProperties(headingLevel string, noOfOccurrences int) {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	modalManager.pageInfoModal.HeadingProperties = append(
		modalManager.pageInfoModal.HeadingProperties,
		Property{
			PropertyName:        headingLevel,
			NumberOfOccurrences: noOfOccurrences,
		},
	)
}

func (modalManager *PageInfoModalManager) GetPageInfoModal() PageInfoModal {
	modalManager.lock.Lock()
	defer modalManager.lock.Unlock()
	return modalManager.pageInfoModal
}
