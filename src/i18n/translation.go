package i18n

type Translation struct {
	MkdirLabel string

	UploadFilesLabel       string
	UploadFilesHint        string
	UploadDirLabel         string
	UploadDirHint          string
	UploadDirContentsLabel string
	UploadDirContentsHint  string
	UploadLabel            string
	UploadingLabel         string

	ListDirLabel  string
	ListNameLabel string
	ListTypeLabel string
	ListSizeLabel string
	ListTimeLabel string

	FilterLabel string

	DeleteConfirm string

	Error403 string
	Error404 string
	Error500 string
}

var translationEnUs = Translation{
	MkdirLabel: "Create dir",

	UploadFilesLabel:       "Files",
	UploadFilesHint:        "Upload files",
	UploadDirLabel:         "Dir",
	UploadDirHint:          "Upload Directory itself",
	UploadDirContentsLabel: "Dir contents",
	UploadDirContentsHint:  "Upload contents of directory",
	UploadLabel:            "Upload",
	UploadingLabel:         "Uploading...",

	ListDirLabel:  "Dir",
	ListNameLabel: "Name",
	ListTypeLabel: "Type",
	ListSizeLabel: "Size",
	ListTimeLabel: "Time",

	FilterLabel: "filter...",

	DeleteConfirm: "Delete?",

	Error403: "403 resource is forbidden",
	Error404: "404 resource not found",
	Error500: "500 potential issue occurred",
}

var translationZhSimp = Translation{
	MkdirLabel: "建目录",

	UploadFilesLabel:       "文件",
	UploadFilesHint:        "上传文件",
	UploadDirLabel:         "目录",
	UploadDirHint:          "上传目录自身",
	UploadDirContentsLabel: "目录内容",
	UploadDirContentsHint:  "上传目录下的内容",
	UploadLabel:            "上传",
	UploadingLabel:         "上传中……",

	ListDirLabel:  "目录",
	ListNameLabel: "名称",
	ListTypeLabel: "类型",
	ListSizeLabel: "大小",
	ListTimeLabel: "时间",

	FilterLabel: "筛选……",

	DeleteConfirm: "确认删除吗？",

	Error403: "403 禁止访问资源",
	Error404: "404 资源不存在",
	Error500: "500 发生潜在错误",
}

var translationZhTrad = Translation{
	MkdirLabel: "建目錄",

	UploadFilesLabel:       "檔案",
	UploadFilesHint:        "上傳檔案",
	UploadDirLabel:         "目錄",
	UploadDirHint:          "上傳目錄自身",
	UploadDirContentsLabel: "目錄內容",
	UploadDirContentsHint:  "上傳目錄下的內容",
	UploadLabel:            "上傳",
	UploadingLabel:         "上傳中……",

	ListDirLabel:  "目錄",
	ListNameLabel: "名稱",
	ListTypeLabel: "類型",
	ListSizeLabel: "大小",
	ListTimeLabel: "時間",

	FilterLabel: "篩選……",

	DeleteConfirm: "確認刪除嗎？",

	Error403: "403 禁止訪問資源",
	Error404: "404 資源不存在",
	Error500: "500 發生潛在錯誤",
}
