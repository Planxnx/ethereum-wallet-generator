package progressbar

import (
	"fmt"

	"github.com/cheggaaa/pb/v3"
	"github.com/schollz/progressbar/v3"
)

const DEFAULT_STANDARD_MODE_TEMPLATE = `{{counters . }} | {{bar . "" "█" "█" "" "" | rndcolor}} | {{percent . }} | {{speed . }} | {{string . "resolved"}}`

//ProgressBar progressbar logging with 2 mode
type ProgressBar struct {
	StandardMode   *pb.ProgressBar
	CompatibleMode *progressbar.ProgressBar
}

//NewCompatibleProgressBar returns a new progress bar set to compatible mode.
func NewCompatibleProgressBar(number int) *ProgressBar {
	return &ProgressBar{
		CompatibleMode: progressbar.NewOptions(number,
			progressbar.OptionSetItsString("w"),
			progressbar.OptionSetPredictTime(true),
			progressbar.OptionShowIts(),
			progressbar.OptionShowCount(),
			progressbar.OptionFullWidth(),
		),
	}
}

//NewStandardProgressBar returns a new progress bar set to standard mode.
func NewStandardProgressBar(number int) *ProgressBar {
	bar := &ProgressBar{
		StandardMode: pb.StartNew(number),
	}
	bar.StandardMode.SetTemplate(pb.Default)
	bar.StandardMode.SetTemplateString(DEFAULT_STANDARD_MODE_TEMPLATE)
	return bar
}

//Increment increment progress
func (bar *ProgressBar) Increment() error {
	if bar.CompatibleMode != nil {
		return bar.CompatibleMode.Add(1)
	}
	return bar.StandardMode.Increment().Err()
}

//SetResolved set resolved wallet number
func (bar *ProgressBar) SetResolved(resolved int) error {
	if bar.StandardMode != nil {
		return bar.StandardMode.Set("resolved", fmt.Sprintf("resolved: %v", resolved)).Err()
	}
	return nil
}

//Finish close progress bar
func (bar *ProgressBar) Finish() error {
	if bar.CompatibleMode != nil {
		return bar.CompatibleMode.Close()
	}
	return bar.StandardMode.Finish().Err()
}
