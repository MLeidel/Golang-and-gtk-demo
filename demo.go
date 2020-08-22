package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

/*
tvtext is set into the textview (txtview) display
*/
var tvtext = `
This is a textview display

with

scrolling

both

horizontal
and
verticle
scrolling.
`
var lbl *gtk.Label // for general use
var str string     // for general use

// // Create global vars for gtk widgets
var g_obj_msgdlg *gtk.MessageDialog
var g_obj_txbuffer *gtk.TextBuffer
var g_obj_lbl_entry *gtk.Label
var g_obj_comboboxoptions *gtk.ComboBoxText
var g_obj_lbl_combo_option *gtk.Label
var g_obj_dlg_about *gtk.AboutDialog
var g_obj_rb_lbl *gtk.Label
var g_obj_dlg_file_open *gtk.FileChooserDialog
var g_obj_entry_file *gtk.Entry
var g_obj_dlg_lbx *gtk.Dialog
var g_obj_listbox *gtk.ListBox

/* struct for checkbox demo */
// use of struct helps clarify
type checks struct {
	chk1 bool
	chk2 bool
	chk3 bool
}

var checks_struct checks

//////////////////////////////////////

func on_listbox_row_activated(oList *gtk.ListBox, oRow *gtk.ListBoxRow) {
	bin, err := oRow.GetChild() // gtk.Bin method
	errorCheck(err)
	lbl = bin.(*gtk.Label)    // type assertion
	txt, err := lbl.GetText() // returns 2
	errorCheck(err)
	g_obj_msgdlg.SetMarkup("<b><big>" + txt + "</big></b>")
	g_obj_msgdlg.FormatSecondaryText("did you click this?")
	g_obj_msgdlg.Run()
	g_obj_msgdlg.Hide()
}

func on_btn_close_clicked(obtn *gtk.Button) {
	os.Exit(0)
}

func on_btn_clear_textiew_clicked() {
	// using global TextBuffer object
	g_obj_txbuffer.SetText("")
}

func on_entry_1_activate(oEntry *gtk.Entry) {
	txt, err := oEntry.GetText() // returns 2
	errorCheck(err)
	g_obj_lbl_entry.SetText(txt)
	oEntry.SetText("")
	if txt == "list" {
		displayListDlg() // launch the dialog with a listbox
	}
}

func displayListDlg() {
	var err error
	var lble *gtk.Label

	// remove any existing rows
	r := g_obj_listbox.GetRowAtIndex(0)
	for r != nil {
		g_obj_listbox.Remove(r)
		r = g_obj_listbox.GetRowAtIndex(0)
	}

	// something to fill the list with
	str = `alpha beta gamma delta iota epsilon
	     lambda zeta kappa`
	rowtext := strings.Fields(str)
	for _, s := range rowtext {
		lble, err = gtk.LabelNew(s)
		errorCheck(err)
		lble.SetXAlign(0.0)
		g_obj_listbox.Insert(lble, 0)
	}

	g_obj_dlg_lbx.ShowAll()
	g_obj_dlg_lbx.Run()
	g_obj_dlg_lbx.Hide()
}

func on_dlg_listbox_row_activated(oList *gtk.ListBox, oRow *gtk.ListBoxRow) {
	on_listbox_row_activated(oList, oRow)
}

func on_cbox_changed(ocbox *gtk.ComboBox) {
	fmt.Println("%d", ocbox.GetActive())
}

func on_btn_show_option_clicked() {
	txt := g_obj_comboboxoptions.GetActiveText()
	g_obj_lbl_combo_option.SetText(txt)
}

func on_btn_about_clicked() {
	g_obj_dlg_about.Run()
	g_obj_dlg_about.Hide()
}

func on_rb_1_toggled() {
	g_obj_rb_lbl.SetText("1%")
}
func on_rb_2_toggled() {
	g_obj_rb_lbl.SetText("2%")
}
func on_rb_3_toggled() {
	g_obj_rb_lbl.SetText("3%")
}

// CheckButtons
// 	another way to do this:
// 	func on_chkbtn_1_toggled(cb* gtk.CheckButton) {
func on_chkbtn_1_toggled() {
	// tb := cb.ToggleButton  // access ToggleButton
	// fmt.Println(tb.GetActive())	// true or false
	if checks_struct.chk1 {
		checks_struct.chk1 = false
	} else {
		checks_struct.chk1 = true
	}
}
func on_chkbtn_2_toggled() {
	if checks_struct.chk2 {
		checks_struct.chk2 = false
	} else {
		checks_struct.chk2 = true
	}
}
func on_chkbtn_3_toggled() {
	if checks_struct.chk3 {
		checks_struct.chk3 = false
	} else {
		checks_struct.chk3 = true
	}
}
func on_btn_check_boxes_clicked() {
	fmt.Println(checks_struct)
	str = ""
	if checks_struct.chk1 {
		str += "Floppy "
	}
	if checks_struct.chk2 {
		str += "HDisk "
	}
	if checks_struct.chk3 {
		str += "CDR "
	}
	g_obj_msgdlg.SetMarkup("You checked:")
	g_obj_msgdlg.FormatSecondaryText(str)
	g_obj_msgdlg.Run()
	g_obj_msgdlg.Hide()
}

func on_btn_openfile_clicked() {
	g_obj_dlg_file_open.Run()
	g_obj_dlg_file_open.Hide()
}

func on_btn_dlg_file_ok_clicked() {
	gfc := g_obj_dlg_file_open.FileChooser
	fpath := gfc.GetFilename()
	g_obj_entry_file.SetText(fpath)
	g_obj_dlg_file_open.Hide()
}

func on_btn_dlg_file_cancel_clicked() {
	g_obj_dlg_file_open.Hide()
}

func on_btn_dlg_close_clicked() {
	g_obj_dlg_lbx.Hide()
}

/* MAIN */

func main() {
	// Create a new application. Change the appID string.
	application, err := gtk.ApplicationNew("com.demo", glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect Builder to application activate event
	application.Connect("activate", func() {

		// Get the GtkBuilder UI definition in the glade file.
		builder, err := gtk.BuilderNewFromFile("demo.glade")
		errorCheck(err)

		// // connect signals mapping handlers to callback functions
		// // NOTE: Most of these callbacks take only 1 parameter
		// // the calling object, even though Glade may allow more than 1.
		// // id : func,
		signals := map[string]interface{}{
			"on_listbox_row_activated":           on_listbox_row_activated,
			"on_btn_close_clicked":               on_btn_close_clicked,
			"on_btn_clear_textiew_clicked":       on_btn_clear_textiew_clicked,
			"on_entry_1_activate":                on_entry_1_activate,
			"on_btn_show_option_clicked":         on_btn_show_option_clicked,
			"on_btn_about_clicked":               on_btn_about_clicked,
			"on_rb_1_toggled":                    on_rb_1_toggled,
			"on_rb_2_toggled":                    on_rb_2_toggled,
			"on_rb_3_toggled":                    on_rb_3_toggled,
			"on_chkbtn_1_toggled":                on_chkbtn_1_toggled,
			"on_chkbtn_2_toggled":                on_chkbtn_2_toggled,
			"on_chkbtn_3_toggled":                on_chkbtn_3_toggled,
			"on_btn_check_boxes_clicked":         on_btn_check_boxes_clicked,
			"on_btn_openfile_clicked":            on_btn_openfile_clicked,
			"on_dlg_file_chooser_file_activated": on_btn_dlg_file_ok_clicked,
			"on_btn_dlg_file_ok_clicked":         on_btn_dlg_file_ok_clicked,
			"on_btn_dlg_file_cancel_clicked":     on_btn_dlg_file_cancel_clicked,
			"on_dlg_listbox_row_activated":       on_dlg_listbox_row_activated,
			"on_btn_dlg_close_clicked":           on_btn_dlg_close_clicked,
		}
		builder.ConnectSignals(signals)

		/*
		   Set local/global access to specific widgets
		   obj_XXXX, err := builder.GetObject("GladeWidget_ID")
		   errorCheck(err)
		   g_obj_XXXX = obj.(*gtk.WidgetName) // assign with type assertion
		*/

		obj_lbl_entry, err := builder.GetObject("lbl_entry_text")
		errorCheck(err)
		g_obj_lbl_entry = obj_lbl_entry.(*gtk.Label)

		// Entry field to hold selected fullpath from FileChooser
		obj_entry_file, err := builder.GetObject("entry_filepath")
		errorCheck(err)
		g_obj_entry_file = obj_entry_file.(*gtk.Entry)

		/* TEXTVIEW */
		obj_tv, err := builder.GetObject("txtview")
		errorCheck(err)
		g_obj_tv := obj_tv.(*gtk.TextView)
		// put text into the TextView/TextBuffer
		buffer, err := g_obj_tv.GetBuffer()
		errorCheck(err)
		g_obj_txbuffer = buffer // global access to this TextBuffer
		g_obj_txbuffer.SetText(tvtext)

		/* LISTBOX */

		str = `alpha beta gamma delta iota epsilon
             lambda zeta kappa`
		rowtext := strings.Fields(str)
		// populate listbox
		obj_lstbox, err := builder.GetObject("listbox")
		errorCheck(err)
		g_obj_lstbox := obj_lstbox.(*gtk.ListBox)
		for _, s := range rowtext {
			lbl, err = gtk.LabelNew(s)
			errorCheck(err)
			g_obj_lstbox.Insert(lbl, 0)
		}

		/* MESSAGEBOX */

		// set a global object to use the MessageBox
		obj_msgdlg, err := builder.GetObject("msgdlg")
		errorCheck(err)
		g_obj_msgdlg = obj_msgdlg.(*gtk.MessageDialog)

		/* COMBOBOXTEXT */

		obj_comboboxoptions, err := builder.GetObject("comboboxoptions")
		errorCheck(err)
		g_obj_comboboxoptions = obj_comboboxoptions.(*gtk.ComboBoxText)
		obj_lbl_combo_option, err := builder.GetObject("lbl_combo_option")
		errorCheck(err)
		g_obj_lbl_combo_option = obj_lbl_combo_option.(*gtk.Label)

		/* ABOUT DIALOG */

		obj_dlg_about, err := builder.GetObject("dlg_about")
		errorCheck(err)
		g_obj_dlg_about = obj_dlg_about.(*gtk.AboutDialog)

		/* RADIOBUTTONS */

		obj_rb_lbl, err := builder.GetObject("lbl_choice")
		errorCheck(err)
		g_obj_rb_lbl = obj_rb_lbl.(*gtk.Label)

		/* CHECKBOX gtk.CheckButton */

		// initialize to all unchecked
		checks_struct.chk1 = false
		checks_struct.chk2 = false
		checks_struct.chk3 = false

		/* FILECHOOSER */

		obj_dlg_file_open, err := builder.GetObject("dlg_file_chooser")
		errorCheck(err)
		g_obj_dlg_file_open = obj_dlg_file_open.(*gtk.FileChooserDialog)

		/* DIALOG GENERAL */

		obj_dlg_lbx, err := builder.GetObject("dialog_list")
		errorCheck(err)
		g_obj_dlg_lbx = obj_dlg_lbx.(*gtk.Dialog)

		obj_listbox, err := builder.GetObject("dlg_listbox")
		errorCheck(err)
		g_obj_listbox = obj_listbox.(*gtk.ListBox)

		/////////////////////////////////////////////
		// show the window object with all widgets //
		/////////////////////////////////////////////
		obj, err := builder.GetObject("window1")
		wnd := obj.(*gtk.Window)
		wnd.ShowAll()
		application.AddWindow(wnd)
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		panic(e)
	}
}
