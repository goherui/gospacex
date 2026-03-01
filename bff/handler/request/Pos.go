package request

type PosCreate struct {
	Title        string `form:"title" json:"title" xml:"title"  binding:"required"`
	Company      string `form:"company" json:"company" xml:"company"  binding:"required"`
	Salary       string `form:"salary" json:"salary" xml:"salary"  binding:"required"`
	Location     string `form:"location" json:"location" xml:"location"  binding:"required"`
	Description  string `form:"description" json:"description" xml:"description"  binding:"required"`
	Requirements string `form:"requirements" json:"requirements" xml:"requirements"  binding:"required"`
}
type PosDel struct {
	Id int `form:"id" json:"id" xml:"id"  binding:"required"`
}
type PosUpdate struct {
	Id           int    `form:"id" json:"id" xml:"id"  binding:"required"`
	Title        string `form:"title" json:"title" xml:"title"  binding:"required"`
	Company      string `form:"company" json:"company" xml:"company"  binding:"required"`
	Salary       string `form:"salary" json:"salary" xml:"salary"  binding:"required"`
	Location     string `form:"location" json:"location" xml:"location"  binding:"required"`
	Description  string `form:"description" json:"description" xml:"description"  binding:"required"`
	Requirements string `form:"requirements" json:"requirements" xml:"requirements"  binding:"required"`
}
type PosList struct {
	Page int `form:"page" json:"page" xml:"page"`
	Size int `form:"size" json:"size" xml:"size"`
}
