package bot

/*
import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

type line struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func (a *Bot) zoneMapStatus(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	type Content struct {
		Message     string
		Status      string
		Runtime     string
		LastStarted time.Time
	}

	var bot *Status
	if bot, err = a.getStatus("zonemap"); err != nil {
				return
	}

	content := &Content{
		Message:     fmt.Sprintf("Bot is %s, last started %s", bot.State, bot.StartTime),
		Status:      bot.State,
		Runtime:     fmt.Sprintf("%.2f minutes", bot.getRuntime().Minutes()),
		LastStarted: bot.StartTime,
	}

		return
}

func (a *Bot) zoneMapCreate(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	type Content struct {
		Message string
	}

	content := &Content{
		Message: "Starting bot to process zonemap",
	}

	if err = a.startBot("zonemap"); err != nil {
				return
	}

	go a.CreateZoneMapCache()

		return
}

//CreateZoneMapCache processes map files and saves them to www/images/maps/ then update zonelevel
func (a *Bot) CreateZoneMapCache() (err error) {
	start := time.Now()
	dir := "map"
	if err = filepath.Walk(dir, a.browseMaps); err != nil {
		return
	}

	a.endBot("zonemap")
	fmt.Println("Finished in", time.Since(start))
	return
}

func (a *Bot) browseMaps(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if !strings.Contains(info.Name(), ".txt") {
		return nil
	}

	err = a.loadMap(path, info.Name())
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Failed to load file %s", info.Name()))
		return err
	}
	return nil
}

func (a *Bot) loadMap(path string, filename string) (err error) {

	if strings.Contains(filename, "_2.txt") {
		return
	}

	bMap, err := ioutil.ReadFile(path)
	if err != nil {
		err = errors.Wrap(err, "Failed to read file")
		return
	}

	r := csv.NewReader(strings.NewReader(string(bMap)))

	//outData := ""
	//iterate each entry
	width := float64(300)
	height := float64(300)
	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.SetLineWidth(0.25)

	var bounds line

	lines := []line{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		if len(record) < 1 {
			continue
		}
		entries := strings.Split(record[0], " ")
		if len(entries) < 1 {
			continue
		}
		drawType := entries[0]
		if drawType != "L" {
			continue
		}

		line := line{}
		line.x1, err = strconv.ParseFloat(strings.TrimSpace(entries[1]), 64) //entries[1]), 64)
		line.y1, err = strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		line.x2, err = strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		line.y2, err = strconv.ParseFloat(strings.TrimSpace(record[4]), 64)
		//line.x1 /= 5
		//line.y1 /= 5
		//line.x2 /= 5
		//line.y2 /= 5

		if bounds.x1 > line.x1 {
			bounds.x1 = line.x1
		}
		if bounds.x2 < line.x1 {
			bounds.x2 = line.x1
		}
		if bounds.x1 > line.x2 {
			bounds.x1 = line.x2
		}
		if bounds.x2 < line.x2 {
			bounds.x2 = line.x2
		}

		if bounds.y1 > line.y1 {
			bounds.y1 = line.y1
		}
		if bounds.y2 < line.y1 {
			bounds.y2 = line.y1
		}
		if bounds.y1 > line.y2 {
			bounds.y1 = line.y2
		}
		if bounds.y2 < line.y2 {
			bounds.y2 = line.y2
		}
		lines = append(lines, line)
		//log.Println("Drew line", line.x1, line.y1, line.x2, line.y2)
		//break
		//outData += fmt.Sprintf("svg.line(g, %f, %f, %f, %f);\n", line.x1, line.y1, line.x2, line.y2)
	}

	var aspect float64

	xOffset := -bounds.x1
	yOffset := -bounds.y1

	bounds.x1 += xOffset
	bounds.x2 += xOffset
	bounds.y1 += yOffset
	bounds.y2 += yOffset

	farPoint := bounds.x2
	if bounds.y2 > farPoint {
		farPoint = bounds.y2
	}

	longestSide := height
	if width > longestSide {
		longestSide = width
	}

	if longestSide > farPoint {
		aspect = farPoint / longestSide
	} else {
		aspect = longestSide / farPoint
	}

	for _, srcLine := range lines {
		line := srcLine
		line.x1 += xOffset
		line.x2 += xOffset
		line.y1 += yOffset
		line.y2 += yOffset

		if bounds.x2 > width {
			line.x1 *= aspect
			line.x2 *= aspect
		}
		if bounds.y2 > height {
			line.y1 *= aspect
			line.y2 *= aspect
		}
		gc.MoveTo(line.x1, line.y1)
		gc.LineTo(line.x2, line.y2)
		gc.Close()
		gc.FillStroke()
	}

	filename = strings.Replace(filename, ".txt", "", -1)
	filename = strings.Replace(filename, "_1", "", -1)
	//log.Println("Aspect ratio is ", aspect, filename)

	if err = draw2dimg.SaveToPngFile("www/images/maps/"+filename+".png", dest); err != nil {
		err = errors.Wrap(err, "Failed to write file")
		return
	}

	zone, err := a.zoneRepo.GetByShortname(filename, nil)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Failed to get %s", filename))
		return
	}
	if zone == nil {
		return
	}

	zoneLevel := &model.ZoneLevel{
		ZoneID:     zone.ID,
		Levels:     zone.Levels,
		MapAspect:  aspect,
		MapXOffset: xOffset,
		MapYOffset: yOffset,
	}

	if err = a.zoneLevelRepo.Edit(zone.ID, zoneLevel, nil); err != nil {
		fmt.Printf("Failed to update %s: %s\n", filename, err.Error())
		err = nil
		//err = errors.Wrap(err, fmt.Sprintf("Failed to update %s", filename))
	}
	//ioutil.WriteFile("out.txt", []byte(outData), 0755)
	return
}
*/
