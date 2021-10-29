---
title: Semantic UI Audit
date: 2021-10-28
---

## Summary

Part of our technical debt wishlist includes replacing Semantic UI's React components with our own, custom-tailored reusable components. In order to prepare for this task, we need to identify what Semantic UI components are currently in use in the Commons client repository.

Therefore, the purpose of this study is to identify:

1. Which Semantic UI components we are using
1. Where are we using them
1. How frequently each one is used

Knowing these data points should help us scope/plan our efforts to remove Semantic UI.

## Rationale

Semantic UI is a very opinionated style framework that strives for visual consistency above all else. While it was great for quickly bootstrapping our UI, as the codebase has grown it has now become more encumbering than helpful. Removing Semantic UI from the codebase would result in several benefits:

- Certain Semantic UI components are not very accessible. We can write our own versions of these components with accessibility as a first-order concern.
- In many places we override Semantic's styling leading to needlessly complicated and heavy-handed CSS. Eliminating Semantic will allow us to reduce global styles, go all-in on CSS modules, and simplify our CSS.
- Semantic UI components are intended to be widely generalizable meaning they often lack functionality we need or add complexity that we don't want. Purpose built components will allow use to create simpler and yet more useful components for our use cases.
- Removing Semantic UI reduces the dependency count for the project. While not a major consideration, pruning our dependency tree is almost always desireable.

By the numbers:

- A total of **150 Commons components** utilize at least one Semantic UI component
- The codebase utilizes **32 distinct Semantic UI components** (as well as 39 sub-components)
- Of these components, **12 are used only once** (likewise, 12 sub-components are used only once)
- Of the **seventeen** instances of `!important` in our CSS, most are used to override Semantic UI

## Semantic Components in Use

The below table indicates which Semantic UI components are currently in use in the Commons Client codebase. The "Instances" column enumerates how many Commons components import a given Semantic UI component. The "Replacement" column indicates which built-in component should be used instead (if such a suitable replacement exists). The final column identifies the status of the Semantic UI component with ✅ indicating that it is no longer in use and ❌ signifying that it is still present in the codebase.

| Semantic UI Component | Instances | Replacement | Status            |
| --------------------- | --------- | ----------- | ----------------- |
| Accordion             | 1         | none        | ❌ - still in use |
| Button                | 41        | none        | ❌ - still in use |
| Card                  | 5         | none        | ❌ - still in use |
| CheckBox              | 5         | none        | ❌ - still in use |
| Confirm               | 7         | none        | ❌ - still in use |
| Container             | 1         | none        | ❌ - still in use |
| Dimmer                | 4         | none        | ❌ - still in use |
| Dropdown              | 5         | none        | ❌ - still in use |
| Embed                 | 4         | none        | ❌ - still in use |
| Form                  | 38        | none        | ❌ - still in use |
| Grid                  | 21        | none        | ❌ - still in use |
| Header                | 5         | none        | ❌ - still in use |
| Icon                  | 20        | none        | ❌ - still in use |
| Image                 | 2         | none        | ❌ - still in use |
| Input                 | 4         | none        | ❌ - still in use |
| Item                  | 1         | none        | ❌ - still in use |
| Label                 | 1         | none        | ❌ - still in use |
| List                  | 10        | none        | ❌ - still in use |
| Loader                | 22        | none        | ❌ - still in use |
| Menu                  | 1         | none        | ❌ - still in use |
| Message               | 1         | none        | ❌ - still in use |
| Modal                 | 20        | none        | ❌ - still in use |
| Pagination            | 2         | none        | ❌ - still in use |
| Placeholder           | 1         | none        | ❌ - still in use |
| Popup                 | 14        | none        | ❌ - still in use |
| Progress              | 1         | none        | ❌ - still in use |
| Segment               | 1         | none        | ❌ - still in use |
| Select                | 1         | none        | ❌ - still in use |
| Step                  | 1         | none        | ❌ - still in use |
| Tab                   | 4         | none        | ❌ - still in use |
| Table                 | 5         | none        | ❌ - still in use |
| TextArea              | 1         | none        | ❌ - still in use |

It should be noted that several of the listed components contain sub-components. In most of these cases, it is actually the sub-components that are of interest to us. As such, when replacing these components, we will need to account for their sub-components as well. The table below lists all components for which we use sub-components and how many time we utilize a given sub-component. This count distinguishes between individual instances of the sub-component and files (i.e. Commons components) that make use of it. In other words, some sub-components are used multiple times in a single Commons component.

| Component   | Sub-Components        | Instances      |
| ----------- | --------------------- | -------------- |
| Button      | Button.Group          | 1 in 1 file    |
| Card        | Card.Content          | 3 in 3 files   |
|             | Card.Group            | 2 in 2 files   |
|             | Card.Header           | 3 in 3 files   |
|             | Card.Meta             | 5 in 3 files   |
| Dimmer      | Dimmer.Dimmable       | 2 in 2 files   |
| Form        | Form.Checkbox         | 4 in 4 files   |
|             | Form.Dropdown         | 19 in 19 files |
|             | Form.Field            | 24 in 11 files |
|             | Form.Group            | 7 in 3 files   |
|             | Form.Input            | 17 in 9 files  |
|             | Form.Radio            | 2 in 2 files   |
|             | Form.TextArea         | 2 in 2 files   |
| Grid        | Grid.Column           | 71 in 21 files |
|             | Grid.Row              | 27 in 16 files |
| Header      | Header.Subheader      | 3 in 1 file    |
| Item        | Item.Content          | 1 in 1 file    |
|             | Item.Group            | 1 in 1 file    |
|             | Item.Header           | 1 in 1 file    |
|             | Item.Image            | 1 in 1 file    |
| List        | List.Content          | 2 in 2 files   |
|             | List.Description      | 2 in 2 files   |
|             | List.Header           | 1 in 1 file    |
|             | List.Icon             | 2 in 2 files   |
|             | List.Item             | 18 in 7 files  |
| Menu        | Menu.Item             | 1 in 1 file    |
| Modal       | Modal.Actions         | 4 in 4 files   |
|             | Modal.Content         | 21 in 17 files |
|             | Modal.Description     | 1 in 1 file    |
| Placeholder | Placeholder.Header    | 1 in 1 file    |
|             | Placeholder.Line      | 12 in 1 file   |
|             | Placeholder.Paragraph | 2 in 1 file    |
| Step        | Step.Group            | 1 in 1 file    |
| Tab         | Tab.Pane              | 8 in 4 files   |
| Table       | Table.Body            | 2 in 2 files   |
|             | Table.Cell            | 2 in 2 files   |
|             | Table.Header          | 1 in 1 file    |
|             | Table.HeaderCell      | 1 in 1 file    |
|             | Table.Row             | 3 in 3 files   |

## Components Using Semantic UI

Below is a table of all Commons components that import at least one Semantic UI component. Given that component names alone can be a bit confusing, we've also included the directory where the given Commons component resides.

| Component                    | Semantic Component(s)               | Path (to parent directory)                                                                               |
| ---------------------------- | ----------------------------------- | -------------------------------------------------------------------------------------------------------- |
| ActionButtons                | Button, Confirm, Modal              | admin/ActionButtons/                                                                                     |
| ButtonPublish                | Button                              | admin/ButtonPublish/                                                                                     |
| CollapsibleSection           | Icon                                | admin/CollapsibleSection/                                                                                |
| Dashboard                    | Grid, Image, Popup, Tab             | admin/Dashboard/                                                                                         |
| DetailsPopup                 | Popup                               | admin/Dashboard/TeamProjects/DetailsPopup/                                                               |
| DetailsPopup                 | Popup                               | admin/Dashboard/MyProjects/DetailsPopup/                                                                 |
| MyProjectPrimaryCol          | Checkbox, Icon, Modal, Popup        | admin/Dashboard/MyProjects/MyProjectPrimaryCol/                                                          |
| ShareProjectItem             | Popup                               | admin/Dashboard/MyProjects/ShareProjectItem/                                                             |
| ShareProjectItem             | Popup                               | admin/Dashboard/TeamProjects/ShareProjectItem/                                                           |
| TeamProjectPrimaryCol        | Checkbox, Icon, Modal, Popup        | admin/Dashboard/TeamProjects/TeamProjectPrimaryCol/                                                      |
| DownloadCaption              | Loader                              | admin/download/DownloadCaption/                                                                          |
| DownloadOtherFiles           | Loader                              | admin/download/DownloadOtherFiles/                                                                       |
| DownloadThumbnail            | Loader                              | admin/download/DownloadThumbnail/                                                                        |
| BureauOfficesDropdown        | Form                                | admin/dropdowns/BureauOfficesDropdown/                                                                   |
| CategoryDropdown             | Form                                | admin/dropdowns/CategoryDropdown/                                                                        |
| CopyrightDropdown            | Form                                | admin/dropdowns/CopyrightDropdown/                                                                       |
| CountriesRegionsDropdown     | Form                                | admin/dropdowns/CountriesRegionsDropdown/                                                                |
| GraphicStyleDropdown         | Form                                | admin/dropdowns/GraphicStyleDropdown/                                                                    |
| LanguageDropdown             | Form                                | admin/dropdowns/LanguageDropdown/                                                                        |
| PackageTypeDropdown          | Form                                | admin/dropdowns/PackageTypeDropdown/                                                                     |
| PolicyPriorityDropdown       | Form                                | admin/dropdowns/PolicyPriorityDropdown/                                                                  |
| QualityDropdown              | Form                                | admin/dropdowns/QualityDropdown/                                                                         |
| SocialPlatformDropdown       | Form                                | admin/dropdowns/SocialPlatformDropdown/                                                                  |
| TagDropdown                  | Form                                | admin/dropdowns/TagDropdown/                                                                             |
| TeamDropdown                 | Form                                | admin/dropdowns/TeamDropdown/                                                                            |
| UseDropdown                  | Form                                | admin/dropdowns/UseDropdown/                                                                             |
| UserDropdown                 | Form                                | admin/dropdowns/UserDropdown/                                                                            |
| VideoBurnedInStatusDropdown  | Form                                | admin/dropdowns/VideoBurnedInStatusDropdown/                                                             |
| VisibilityDropdown           | Form                                | admin/dropdowns/VisibilityDropdown/                                                                      |
| DynamicConfirm               | Confirm                             | admin/DynamicConfirm/                                                                                    |
| EditFileGrid                 | Button                              | admin/EditFileGrid/                                                                                      |
| EditFileGridRow              | Button, Icon                        | admin/EditFileGridRow/                                                                                   |
| FileRemoveReplaceButtonGroup | Button                              | admin/FileRemoveReplaceButtonGroup/                                                                      |
| FileRemoveReplaceMenu        | Button, Popup                       | admin/FileRemoveReplaceMenu/                                                                             |
| FileUploadProgressBar        | Progress                            | admin/FileUploadProgressBar/                                                                             |
| PackageEdit                  | Button, Loader                      | admin/PackageEdit/                                                                                       |
| EditPackageFilesModal        | Button, Dimmer, Form, Header, Modal | admin/PackageEdit/EditPackageFilesModal/                                                                 |
| EditPressOfficeFileRow       | Grid                                | admin/PackageEdit/EditPressOfficeFileRow/                                                                |
| EditPressOfficeFilesGrid     | Grid                                | admin/PackageEdit/EditPressOfficeFilesGrid/                                                              |
| PackageFiles                 | Button                              | admin/PackageEdit/PackageFiles/                                                                          |
| PlaybookPreview              | Loader                              | admin/Previews/PlaybookPreview/                                                                          |
| PreviewLoader                | Loader                              | admin/Previews/PreviewLoader/                                                                            |
| ProjectPreviewContent        | Dropdown, Embed                     | admin/Previews/ProjectPreview/ProjectPreviewContent/                                                     |
| ProjectDetailsForm           | Button, Form, Grid, Input, TextArea | admin/ProjectDetailsForm/                                                                                |
| EditProjectFilesModal        | Button, Form, Dimmer, Header, Modal | admin/ProjectEdit/EditProjectFilesModal/                                                                 |
| EditSupportFileRow           | Grid                                | admin/ProjectEdit/EditProjectFilesModal/EditSupportFileRow/                                              |
| EditSupportFilesGrid         | Grid                                | admin/ProjectEdit/EditProjectFilesModal/EditSupportFilesGrid/                                            |
| EditVideoFileRow             | Grid                                | admin/ProjectEdit/EditProjectFilesModal/EditVideoFileRow/                                                |
| EditVideoFilesGrid           | Grid                                | admin/ProjectEdit/EditProjectFilesModal/EditVideoFilesGrid/                                              |
| EditSingleProjectItem        | Loader                              | admin/ProjectEdit/EditSingleProjectItem/                                                                 |
| Carousel                     | Icon                                | admin/ProjectEdit/EditVideoModal/Carousel/                                                               |
| FileDataForm                 | Confirm, Form, Grid                 | admin/ProjectEdit/EditVideoModal/ModalForms/FileDataForm/                                                |
| UnitDataForm                 | Embed, Form, Grid                   | admin/ProjectEdit/EditVideoModal/ModalForms/UnitDataForm/                                                |
| FileSection                  | Icon                                | admin/ProjectEdit/EditVideoModal/ModalSections/FileSection/                                              |
| GraphicEdit                  | Button, Loader, Modal               | admin/ProjectEdit/GraphicEdit/                                                                           |
| GraphicFilesForm             | Confirm, Form, Input, Loader        | admin/ProjectEdit/GraphicEdit/ GraphicFilesFormContainer/GraphicFilesForm/                               |
| GraphicSupportFiles          | Confirm                             | admin/ProjectEdit/GraphicEdit/GraphicSupportFiles/                                                       |
| withModal                    | Modal                               | admin/ProjectEdit/withModal/                                                                             |
| ProjectHeader                | Icon                                | admin/ProjectHeader/                                                                                     |
| VideoProjectData             | Icon, Loader                        | admin/ProjectReview/VideoProjectData/                                                                    |
| VideoProjectFiles            | Button, Grid, Loader                | admin/ProjectReview/VideoProjectFiles/                                                                   |
| VideoProjectFile             | Embed, Grid                         | admin/ProjectReview/VideoProjectFiles/VideoProjectFile/                                                  |
| VideoReview                  | Button, Confirm, Grid, Icon, Loader | admin/ProjectReview/VideoReview/                                                                         |
| VideoSupportFiles            | Icon, Loader                        | admin/ProjectReview/VideoSupportFiles/                                                                   |
| ProjectSupportFiles          | Grid                                | admin/ProjectSupportFiles/                                                                               |
| SupportItem                  | Loader, Popup                       | admin/ProjectSupportFiles/SupportItem                                                                    |
| ProjectUnits                 | Card                                | admin/ProjectUnits/                                                                                      |
| ProjectUnitItem              | Card, Image, List, Loader, Modal    | admin/ProjectUnits/ProjectUnitItem/                                                                      |
| TermsConditions              | Form                                | admin/TermsConditions/                                                                                   |
| Upload                       | Button, Icon, Modal                 | admin/Upload/                                                                                            |
| CancelUpload                 | Button, Modal                       | admin/Upload/modals/CancelUpload/                                                                        |
| Confirm                      | Confirm                             | admin/Upload/modals/Confirm/                                                                             |
| IncludeRequiredFileMsg       | Modal                               | admin/Upload/modals/IncludeRequiredFileMsg/                                                              |
| VideoUpload                  | Dimmer, Loader, Tab                 | admin/Upload/modals/VideoUpload/                                                                         |
| UploadCompletionTracker      | Label, Icon                         | admin/Upload/modals/VideoUpload/ VideoProjectFiles/UploadCompletionTracker/                              |
| VideoProjectFilesDesktop     | Button, Form, Grid, Step            | admin/Upload/modals/VideoUpload/ VideoProjectFiles/VideoProjectFilesDesktop/                             |
| VideoProjectFilesDesktopRow  | Grid                                | admin/Upload/modals/VideoUpload/VideoProjectFiles/ VideoProjectFilesDesktop/VideoProjectFilesDesktopRow/ |
| VideoProjectFilesMobile      | Button, Form                        | admin/Upload/modals/VideoUpload/VideoProjectFiles/ VideoProjectFilesMobile/                              |
| VideoProjectFilesRowMobile   | Button                              | admin/Upload/modals/VideoUpload/VideoProjectFiles/ VideoProjectFilesMobile/VideoProjectFilesRowMobile/   |
| VideoProjectType             | Button, Form                        | admin/Upload/modals/VideoUpload/VideoProjectType/                                                        |
| UploadSuccessMsg             | Icon                                | admin/UploadSuccessMsg/                                                                                  |
| ButtonAddFiles               | Button                              | ButtonAddFiles/                                                                                          |
| ClipboardCopy                | Button                              | ClipboardCopy/                                                                                           |
| Document                     | Button                              | Document/                                                                                                |
| DocumentCard                 | Card, Modal                         | Document/DocumentCard/                                                                                   |
| DocumentationMenu            | List                                | DocumentationMenu/                                                                                       |
| DocumentationSidebar         | List                                | DocumentationSidebar/                                                                                    |
| EmailRequest                 | Button, Form                        | EmailRequest/                                                                                            |
| FeaturedError                | Message                             | Featured/FeaturedError/                                                                                  |
| FeaturedLoading              | Loader                              | Featured/FeaturedLoading/                                                                                |
| Packages                     | Grid                                | Featured/Packages/                                                                                       |
| Priorities                   | Modal                               | Featured/Priorities/                                                                                     |
| Recents                      | Modal                               | Featured/Recents/                                                                                        |
| FileListDisplay              | Dimmer, Loader, Segment             | FileListDisplay/                                                                                         |
| FilterMenuCountries          | Form                                | FilterMenu/FilterMenuCountries/                                                                          |
| FilterMenuItem               | Form, Icon                          | FilterMenu/FilterMenuItem/                                                                               |
| Footer                       | Container, Header, List             | Footer/                                                                                                  |
| GenericLoader                | Loader                              | GenericLoader/                                                                                           |
| GraphicCard                  | Modal                               | GraphicProject/GraphicCard/                                                                              |
| EmailLogin                   | Button, Form                        | Login/EmailLogin/                                                                                        |
| GoogleLogin                  | Button                              | Login/GoogleLogin/                                                                                       |
| Login                        | Button                              | Login/                                                                                                   |
| Logout                       | Button                              | Logout/                                                                                                  |
| ModalContentMeta             | Icon                                | modals/ModalContentMeta/                                                                                 |
| ModalLangDropdown            | Dropdown                            | modals/ModalLangDropdown/                                                                                |
| GlobalNav                    | Loader                              | navs/global/                                                                                             |
| LoggedInNav                  | Popup                               | navs/global/LoggedInNav/                                                                                 |
| Notification                 | Icon                                | Notification/                                                                                            |
| Package                      | Card                                | Package/                                                                                                 |
| PackageCard                  | Card, Modal                         | Package/PackageCard                                                                                      |
| PasswordReset                | Button, Form                        | PasswordReset/                                                                                           |
| DocumentPlaceholder          | Placeholder                         | Placeholder/DocumentPlaceholder/                                                                         |
| Popup                        | Header                              | popups/Popup/                                                                                            |
| PopupTabbed                  | Header, Tab                         | popups/PopupTabbed/                                                                                      |
| PopupTrigger                 | Button, Popup                       | popups/PopupTrigger/                                                                                     |
| IconPopup                    | Icon, Popup                         | popups/IconPopup/                                                                                        |
| Register                     | Tab                                 | Register/                                                                                                |
| ReviewSubmit                 | Button, Form, List                  | Register/ReviewSubmit/                                                                                   |
| SelectRole                   | Button, Form, Input                 | Register/SelectRole/                                                                                     |
| TeamDetails                  | Button, Form                        | Register/TeamDetails/                                                                                    |
| UserDetails                  | Button, Form                        | Register/UserDetails/                                                                                    |
| RegisterConfirmation         | Button, Form                        | RegisterConfirmation/                                                                                    |
| Results                      | Grid                                | Results/                                                                                                 |
| ResultItem                   | Modal                               | Results/ResultItem/                                                                                      |
| ResultHeader                 | Dropdown                            | Results/ResultHeader/                                                                                    |
| ResultsPagination            | Pagination                          | Results/ResultsPagination/                                                                               |
| ResultsToggleView            | Icon                                | Results/ResultsToggleView/                                                                               |
| ScrollableTableWithMenu      | Grid, Table                         | ScrollableTableWithMenu/                                                                                 |
| TableActionsMenu             | Button, Checkbox, Modal             | ScrollableTableWithMenu/TableActionsMenu/                                                                |
| ActionResults                | List                                | ScrollableTableWithMenu/TableActionsMenu/ActionResults/                                                  |
| ActionResultsError           | List                                | ScrollableTableWithMenu/TableActionsMenu/ ActionResults/ActionResultsError/                              |
| ActionResultsItem            | List                                | ScrollableTableWithMenu/TableActionsMenu/ ActionResults/ActionResultsItem/                               |
| DeleteIconButton             | Button, Popup                       | ScrollableTableWithMenu/TableActionsMenu/DeleteIconButton/                                               |
| DeleteProjects               | Button, Modal                       | ScrollableTableWithMenu/TableActionsMenu/DeleteProjects/                                                 |
| UnpublishProjects            | Button, Popup                       | ScrollableTableWithMenu/TableActionsMenu/UnpublishProjects/                                              |
| TableBody                    | Table                               | ScrollableTableWithMenu/TableBody/                                                                       |
| TableBodyMessage             | Loader, Table                       | ScrollableTableWithMenu/TableBody/TableBodyMessage/                                                      |
| TableHeader                  | Table                               | ScrollableTableWithMenu/TableHeader/                                                                     |
| TableItemsDisplay            | Dropdown, Grid, Loader              | ScrollableTableWithMenu/TableItemsDisplay/                                                               |
| TableMenu                    | Accordion, Checkbox, Icon, Menu     | ScrollableTableWithMenu/TableMenu/                                                                       |
| TableMobileDataToggleIcon    | Icon                                | ScrollableTableWithMenu/TableMobileDataToggleIcon/                                                       |
| TablePagination              | Pagination                          | ScrollableTableWithMenu/TablePagination/                                                                 |
| TableRow                     | Table                               | ScrollableTableWithMenu/TableRow/                                                                        |
| TableSearch                  | Button, Form, Grid, Input           | ScrollableTableWithMenu/TableSearch/                                                                     |
| SearchInput                  | Dropdown                            | SearchInput/                                                                                             |
| Share                        | List                                | Share/                                                                                                   |
| ShareButton                  | List                                | Share/ShareButton/                                                                                       |
| UserAdmin                    | Button, Form, Modal, Select         | User/UserAdmin/                                                                                          |
| Video                        | Checkbox, Embed                     | Video/                                                                                                   |
| DownloadItem                 | Item                                | Video/Download/DownloadItem/                                                                             |
