
package river

import (
  "fmt"
  "strings"
  ui "github.com/gizak/termui"
)

var (
  NItems = 0
  ReceiveBoxHeight = 50
)

func CreateReceiveBox() *ui.Par {
  p := ui.NewPar("")
  // p.Height = ui.TermHeight() - SendBoxHeight
  p.Height = ReceiveBoxHeight
  p.TextFgColor = ui.ColorWhite
  p.BorderFg = ui.ColorCyan
  go ReadReceives(p)
  return p
}

func ReadReceives(p *ui.Par) {
  for {
    NItems += 1
    if NItems >= ReceiveBoxHeight {
      ScrollReceiveBox(p)
    }
    msg := <-Display
    if msg.Type == SENDING {
      p.Text += fmt.Sprintf(" -> %v\n", msg.Message)
    } else if msg.Type == RECEIVING {
      p.Text += fmt.Sprintf(" <- %v\n", msg.Message)
    } else if msg.Type == ERROR {
      p.Text += fmt.Sprintf(" !! %v\n", msg.Message)
    } else if msg.Type == COMMAND {
      p.Text += fmt.Sprintf(" <> %v\n", msg.Message)
    }
    ui.Render(ui.Body)
  }
}

func ScrollReceiveBox(p *ui.Par) {
  sp := strings.Split(p.Text, "\n")
  p.Text = strings.Join(sp[1:], "\n")
}
