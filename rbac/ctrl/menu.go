package ctrl

type menuCtrl struct {

}
var menu *menuCtrl

func MenuCtrl()*menuCtrl {
	if menu ==nil {
		menu =new(menuCtrl)
	}
	return menu
}

func (*menuCtrl) List() {

}