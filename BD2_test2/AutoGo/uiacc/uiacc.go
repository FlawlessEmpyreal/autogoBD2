package uiacc

type Uiacc struct {
}

type UiObject struct {
}

type Rect struct {
	Left    int
	Right   int
	Top     int
	Bottom  int
	CenterX int
	CenterY int
	Width   int
	Height  int
}

// New 创建一个新的 Accessibility 对象,安卓11开始才支持操作虚拟屏节点
func New(displayId int) *Uiacc {
	return nil
}

// Text 设置选择器的 text 属性
func (a *Uiacc) Text(value string) *Uiacc {
	return nil
}

// TextContains 设置选择器的 textContains 属性，用于匹配包含指定文本的控件
func (a *Uiacc) TextContains(value string) *Uiacc {
	return nil
}

// TextStartsWith 设置选择器的 textStartsWith 属性，用于匹配以指定文本开头的控件
func (a *Uiacc) TextStartsWith(value string) *Uiacc {
	return nil
}

// TextEndsWith 设置选择器的 textEndsWith 属性，用于匹配以指定文本结尾的控件
func (a *Uiacc) TextEndsWith(value string) *Uiacc {
	return nil
}

// TextMatches 设置选择器的 textMatches 属性，用于匹配符合指定正则表达式的控件
func (a *Uiacc) TextMatches(value string) *Uiacc {
	return nil
}

// Desc 设置选择器的 desc 属性，用于匹配描述等于指定文本的控件
func (a *Uiacc) Desc(value string) *Uiacc {
	return nil
}

// DescContains 设置选择器的 descContains 属性，用于匹配描述包含指定文本的控件
func (a *Uiacc) DescContains(value string) *Uiacc {
	return nil
}

// DescStartsWith 设置选择器的 descStartsWith 属性，用于匹配描述以指定文本开头的控件
func (a *Uiacc) DescStartsWith(value string) *Uiacc {
	return nil
}

// DescEndsWith 设置选择器的 descEndsWith 属性，用于匹配描述以指定文本结尾的控件
func (a *Uiacc) DescEndsWith(value string) *Uiacc {
	return nil
}

// DescMatches 设置选择器的 descMatches 属性，用于匹配描述符合指定正则表达式的控件
func (a *Uiacc) DescMatches(value string) *Uiacc {
	return nil
}

// Id 设置选择器的 id 属性，用于匹配ID等于指定值的控件
func (a *Uiacc) Id(value string) *Uiacc {
	return nil
}

// IdContains 设置选择器的 idContains 属性，用于匹配ID包含指定值的控件
func (a *Uiacc) IdContains(value string) *Uiacc {
	return nil
}

// IdStartsWith 设置选择器的 idStartsWith 属性，用于匹配ID以指定值开头的控件
func (a *Uiacc) IdStartsWith(value string) *Uiacc {
	return nil
}

// IdEndsWith 设置选择器的 idEndsWith 属性，用于匹配ID以指定值结尾的控件
func (a *Uiacc) IdEndsWith(value string) *Uiacc {
	return nil
}

// IdMatches 设置选择器的 idMatches 属性，用于匹配ID符合指定正则表达式的控件
func (a *Uiacc) IdMatches(value string) *Uiacc {
	return nil
}

// ClassName 设置选择器的 className 属性，用于匹配类名等于指定值的控件
func (a *Uiacc) ClassName(value string) *Uiacc {
	return nil
}

// ClassNameContains 设置选择器的 classNameContains 属性，用于匹配类名包含指定值的控件
func (a *Uiacc) ClassNameContains(value string) *Uiacc {
	return nil
}

// ClassNameStartsWith 设置选择器的 classNameStartsWith 属性，用于匹配类名以指定值开头的控件
func (a *Uiacc) ClassNameStartsWith(value string) *Uiacc {
	return nil
}

// ClassNameEndsWith 设置选择器的 classNameEndsWith 属性，用于匹配类名以指定值结尾的控件
func (a *Uiacc) ClassNameEndsWith(value string) *Uiacc {
	return nil
}

// ClassNameMatches 设置选择器的 classNameMatches 属性，用于匹配类名符合指定正则表达式的控件
func (a *Uiacc) ClassNameMatches(value string) *Uiacc {
	return nil
}

// PackageName 设置选择器的 packageName 属性，用于匹配包名等于指定值的控件
func (a *Uiacc) PackageName(value string) *Uiacc {
	return nil
}

// PackageNameContains 设置选择器的 packageNameContains 属性，用于匹配包名包含指定值的控件
func (a *Uiacc) PackageNameContains(value string) *Uiacc {
	return nil
}

// PackageNameStartsWith 设置选择器的 packageNameStartsWith 属性，用于匹配包名以指定值开头的控件
func (a *Uiacc) PackageNameStartsWith(value string) *Uiacc {
	return nil
}

// PackageNameEndsWith 设置选择器的 packageNameEndsWith 属性，用于匹配包名以指定值结尾的控件
func (a *Uiacc) PackageNameEndsWith(value string) *Uiacc {
	return nil
}

// PackageNameMatches 设置选择器的 packageNameMatches 属性，用于匹配包名符合指定正则表达式的控件
func (a *Uiacc) PackageNameMatches(value string) *Uiacc {
	return nil
}

// Bounds 设置选择器的 bounds 属性，用于匹配控件在屏幕上的范围
func (a *Uiacc) Bounds(left, top, right, bottom int) *Uiacc {
	return nil
}

// BoundsInside 设置选择器的 boundsInside 属性，用于匹配控件在屏幕内的范围
func (a *Uiacc) BoundsInside(left, top, right, bottom int) *Uiacc {
	return nil
}

// BoundsContains 设置选择器的 boundsContains 属性，用于匹配控件包含在指定范围内
func (a *Uiacc) BoundsContains(left, top, right, bottom int) *Uiacc {
	return nil
}

// DrawingOrder 设置选择器的 drawingOrder 属性，用于匹配控件在父控件中的绘制顺序
func (a *Uiacc) DrawingOrder(value int) *Uiacc {
	return nil
}

// Clickable 设置选择器的 clickable 属性，用于匹配控件是否可点击
func (a *Uiacc) Clickable(value bool) *Uiacc {
	return nil
}

// LongClickable 设置选择器的 longClickable 属性，用于匹配控件是否可长按
func (a *Uiacc) LongClickable(value bool) *Uiacc {
	return nil
}

// Checkable 设置选择器的 checkable 属性，用于匹配控件是否可选中
func (a *Uiacc) Checkable(value bool) *Uiacc {
	return nil
}

// Selected 设置选择器的 selected 属性，用于匹配控件是否被选中
func (a *Uiacc) Selected(value bool) *Uiacc {
	return nil
}

// Enabled 设置选择器的 enabled 属性，用于匹配控件是否启用
func (a *Uiacc) Enabled(value bool) *Uiacc {
	return nil
}

// Scrollable 设置选择器的 scrollable 属性，用于匹配控件是否可滚动
func (a *Uiacc) Scrollable(value bool) *Uiacc {
	return nil
}

// Editable 设置选择器的 editable 属性，用于匹配控件是否可编辑
func (a *Uiacc) Editable(value bool) *Uiacc {
	return nil
}

// MultiLine 设置选择器的 multiLine 属性，用于匹配控件是否多行
func (a *Uiacc) MultiLine(value bool) *Uiacc {
	return nil
}

// Checked 设置选择器的 checked 属性，用于匹配控件是否被勾选
func (a *Uiacc) Checked(value bool) *Uiacc {
	return nil
}

// Focusable 设置选择器的 focusable 属性，用于匹配控件是否可聚焦
func (a *Uiacc) Focusable(value bool) *Uiacc {
	return nil
}

// Dismissable 设置选择器的 dismissable 属性，用于匹配控件是否可解散
func (a *Uiacc) Dismissable(value bool) *Uiacc {
	return nil
}

// Focused 设置选择器的 focused 属性，用于匹配控件是否是辅助功能焦点
func (a *Uiacc) Focused(value bool) *Uiacc {
	return nil
}

// ContextClickable 设置选择器的 contextClickable 属性，用于匹配控件是否是上下文点击
func (a *Uiacc) ContextClickable(value bool) *Uiacc {
	return nil
}

// Index 设置选择器的 index 属性，用于匹配控件在父控件中的索引
func (a *Uiacc) Index(value int) *Uiacc {
	return nil
}

// Visible 设置选择器的 visible 属性，用于匹配控件是否可见
func (a *Uiacc) Visible(value bool) *Uiacc {
	return nil
}

// Password 设置选择器的 password 属性，用于匹配控件是否为密码字段
func (a *Uiacc) Password(value bool) *Uiacc {
	return nil
}

// Click 点击屏幕上的文本
func (a *Uiacc) Click(text string) bool {
	return false
}

// WaitFor 等待控件出现并返回 UiObject 对象 超时单位为毫秒,写0代表无限等待,超时返回nil
func (a *Uiacc) WaitFor(timeout int) *UiObject {
	return nil
}

// FindOnce 查找单个控件并返回 UiObject 对象
func (a *Uiacc) FindOnce() *UiObject {
	return nil
}

// Find 查找所有符合条件的控件并返回 UiObject 对象数组
func (a *Uiacc) Find() []*UiObject {
	return nil
}

// Release 释放无障碍对象
func (a *Uiacc) Release() {

}

// Click 点击该控件，并返回是否点击成功
func (u *UiObject) Click() bool {
	return false
}

// ClickCenter 使用坐标点击该控件的中点，相当于click(uiObj.bounds().centerX(), uiObject.bounds().centerY())
func (u *UiObject) ClickCenter() bool {
	return false
}

// ClickLongClick 长按该控件，并返回是否点击成功
func (u *UiObject) ClickLongClick() bool {
	return false
}

// Copy 对输入框文本的选中内容进行复制，并返回是否操作成功
func (u *UiObject) Copy() bool {
	return false
}

// Cut 对输入框文本的选中内容进行剪切，并返回是否操作成功
func (u *UiObject) Cut() bool {
	return false
}

// Paste 对输入框控件进行粘贴操作，把剪贴板内容粘贴到输入框中，并返回是否操作成功
func (u *UiObject) Paste() bool {
	return false
}

// ScrollForward 对控件执行向前滑动的操作，并返回是否操作成功
func (u *UiObject) ScrollForward() bool {
	return false
}

// ScrollBackward 对控件执行向后滑动的操作，并返回是否操作成功
func (u *UiObject) ScrollBackward() bool {
	return false
}

// Collapse 对控件执行折叠操作，并返回是否操作成功
func (u *UiObject) Collapse() bool {
	return false
}

// Expand 对控件执行展开操作，并返回是否操作成功
func (u *UiObject) Expand() bool {
	return false
}

// Show 执行显示操作，并返回是否操作成功
func (u *UiObject) Show() bool {
	return false
}

// Select 对控件执行"选中"操作，并返回是否操作成功
func (u *UiObject) Select() bool {
	return false
}

// ClearSelect 清除控件的选中状态，并返回是否操作成功
func (u *UiObject) ClearSelect() bool {
	return false
}

// SetSelection 对输入框控件设置选中的文字内容，并返回是否操作成功
func (u *UiObject) SetSelection(start, end int) bool {
	return false
}

// SetVisibleToUser 设置控件是否可见，并返回是否操作成功
func (u *UiObject) SetVisibleToUser(isVisible bool) bool {
	return false
}

// SetText 设置输入框控件的文本内容，并返回是否设置成功
func (u *UiObject) SetText(str string) bool {
	return false
}

// GetClickable 获取控件的 clickable 属性
func (u *UiObject) GetClickable() bool {
	return false
}

// GetLongClickable 获取控件的 longClickable 属性
func (u *UiObject) GetLongClickable() bool {
	return false
}

// GetCheckable 获取控件的 checkable 属性
func (u *UiObject) GetCheckable() bool {
	return false
}

// GetSelected 获取控件的 selected 属性
func (u *UiObject) GetSelected() bool {
	return false
}

// GetEnabled 获取控件的 enabled 属性
func (u *UiObject) GetEnabled() bool {
	return false
}

// GetScrollable 获取控件的 scrollable 属性
func (u *UiObject) GetScrollable() bool {
	return false
}

// GetEditable 获取控件的 editable 属性
func (u *UiObject) GetEditable() bool {
	return false
}

// GetMultiLine 获取控件的 multiLine 属性
func (u *UiObject) GetMultiLine() bool {
	return false
}

// GetChecked 获取控件的 checked 属性
func (u *UiObject) GetChecked() bool {
	return false
}

// GetFocused 获取控件的 focused 属性
func (u *UiObject) GetFocused() bool {
	return false
}

// GetFocusable 获取控件的 focusable 属性
func (u *UiObject) GetFocusable() bool {
	return false
}

// GetDismissable 获取控件的 dismissable 属性
func (u *UiObject) GetDismissable() bool {
	return false
}

// GetContextClickable 获取控件的 contextClickable 属性
func (u *UiObject) GetContextClickable() bool {
	return false
}

// GetVisible 获取控件的 visible 属性
func (u *UiObject) GetVisible() bool {
	return false
}

// GetPassword 获取控件的 password 属性
func (u *UiObject) GetPassword() bool {
	return false
}

// GetAccessibilityFocused 获取控件的 AccessibilityFocused 属性
func (u *UiObject) GetAccessibilityFocused() bool {
	return false
}

// GetChildCount 获取控件的子控件数目
func (u *UiObject) GetChildCount() int {
	return 0
}

// GetDrawingOrder 获取控件在父控件中的绘制次序
func (u *UiObject) GetDrawingOrder() int {
	return 0
}

// GetIndex 获取控件在父控件中的索引
func (u *UiObject) GetIndex() int {
	return 0
}

// GetBounds 获取控件在屏幕上的范围
func (u *UiObject) GetBounds() Rect {
	return Rect{}
}

// GetBoundsInParent 获取控件在父控件中的范围
func (u *UiObject) GetBoundsInParent() Rect {
	return Rect{}
}

// GetId 获取控件的ID
func (u *UiObject) GetId() string {
	return ""
}

// GetText 获取控件的文本内容
func (u *UiObject) GetText() string {
	return ""
}

// GetDesc 获取控件的描述内容
func (u *UiObject) GetDesc() string {
	return ""
}

// GetPackageName 获取控件的包名
func (u *UiObject) GetPackageName() string {
	return ""
}

// GetClassName 获取控件的类名
func (u *UiObject) GetClassName() string {
	return ""
}

// GetParent 获取控件的父控件
func (u *UiObject) GetParent() *UiObject {
	return nil
}

// GetChild 获取控件的指定索引的子控件
func (u *UiObject) GetChild(index int) *UiObject {
	return nil
}

// GetChildren 获取控件的所有子控件
func (u *UiObject) GetChildren() []*UiObject {
	return nil
}

// ToString 节点对象转文本
func (u *UiObject) ToString() string {
	return ""
}
