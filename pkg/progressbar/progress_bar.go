package progressbar

import (
	"fmt"

	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

const DefaultStandardModeTemplate = `{{counters . }} | {{bar . "" "█" "█" "" "" | rndcolor}} | {{percent . }} | {{speed . }} | {{string . "resolved"}}`

// ProgressBar progressbar logging with 2 mode
type ProgressBar struct {
	StandardMode   *pb.ProgressBar
	CompatibleMode *progressbar.ProgressBar
}

// NewCompatibleProgressBar returns a new progress bar set to compatible mode.
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

// NewStandardProgressBar returns a new progress bar set to standard mode.
func NewStandardProgressBar(number int) *ProgressBar {
	bar := &ProgressBar{
		StandardMode: pb.StartNew(number),
	}
	bar.StandardMode.SetTemplate(pb.Default)
	bar.StandardMode.SetTemplateString(DefaultStandardModeTemplate)
	return bar
}

// Increment increment progress
func (bar *ProgressBar) Increment() error {
	if bar.CompatibleMode != nil {
		return errors.WithStack(bar.CompatibleMode.Add(1))
	}
	return errors.WithStack(bar.StandardMode.Increment().Err())
}

// SetResolved set resolved wallet number
func (bar *ProgressBar) SetResolved(resolved int) error {
	if bar.StandardMode != nil {
		return errors.WithStack(bar.StandardMode.Set("resolved", fmt.Sprintf("resolved: %v", resolved)).Err())
	}
	return nil
}

// Finish close progress bar
func (bar *ProgressBar) Finish() error {
	if bar.CompatibleMode != nil {
		return errors.WithStack(bar.CompatibleMode.Close())
	}
	return errors.WithStack(bar.StandardMode.Finish().Err())
}
