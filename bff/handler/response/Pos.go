package response

type PosCreate struct {
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg"  binding:"required"`
}
type PosDel struct {
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg"  binding:"required"`
}
type PosCUpdate struct {
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg"  binding:"required"`
}
type Pos struct {
	ID           int64  `form:"id" json:"id" xml:"id"`
	Title        string `form:"title" json:"title" xml:"title"`
	Company      string `form:"company" json:"company" xml:"company"`
	Salary       string `form:"salary" json:"salary" xml:"salary"`
	Location     string `form:"location" json:"location" xml:"location"`
	Description  string `form:"description" json:"description" xml:"description"`
	Requirements string `form:"requirements" json:"requirements" xml:"requirements"`
}

type PosList struct {
	List []Pos  `form:"list" json:"list" xml:"list"`
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`
}
