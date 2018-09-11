package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"sort"
	"strconv"
	"time"
)

func create(callLog string) {
	filepath := callLog
	xlFile, err := xlsx.OpenFile(filepath)
	if err != nil {
		// create error message -- invalid file
	}
	communities := make(map[string]int)
	person := make(map[string]int)
	personcalls := make(map[string]int)
	personnotes := make(map[string]int)
	communitynotes := make(map[string]int)
	communitycalls := make(map[string]int)
	personcommunity := make(map[string]map[string]map[string]int)

	var name string
	var community string
	var contact string

	for _, sheet := range xlFile.Sheets {
		for rownum, row:= range sheet.Rows {
			if rownum > 0 {
				for cellnum, cell := range row.Cells {
					contents := cell.String()
					if cellnum == 0 {
						community = contents
						if _, ok := communities[community]; ok {
							//fmt.Printf("%s\n", contents + " " + strconv.Itoa(cellnum))
							communities[community] += 1
						} else {
							communities[community] = 1
							communitynotes[community] = 0
							communitycalls[community] = 0
						}
					} else if cellnum == 1 {
						name = contents
						if name == "" {
							name = "NULL"
						}
						if _, ok := person[name]; ok {
							person[name] += 1
						} else {
							person[name] = 1
							personnotes[name] = 0
							personcalls[name] = 0
						}
						if _, ok := personcommunity[community]; ok == false {
							personcommunity[community] = make(map[string]map[string]int)
						}
						if _, ok := personcommunity[community][name]; ok == false {
							personcommunity[community][name] = map[string]int{"call":0, "note":0}
						}
					} else if cellnum == 5 {
						contact = contents
						if contents == "note" {
							personnotes[name] += 1
							communitynotes[community] += 1
						} else if contents == "call" {
							personcalls[name] += 1
							communitycalls[community] += 1
						}
						personcommunity[community][name][contact] += 1
					}
				}
			}
		}
	}

	//this is to write the file
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	//var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	var totalnc int
	for _, v := range person {
		totalnc += v
	}
	var totalnotes int
	for _, v := range personnotes {
		totalnotes += v
	}
	var totalcalls int
	for _, v := range personcalls {
		totalcalls += v
	}

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = ""
	cell = row.AddCell()
	cell.Value = "Total Notes"
	cell = row.AddCell()
	cell.Value = "Total Calls"
	cell = row.AddCell()
	cell.Value = "Notes+Calls"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Total"
	cell = row.AddCell()
	cell.Value = strconv.Itoa(totalnotes)
	cell = row.AddCell()
	cell.Value = strconv.Itoa(totalcalls)
	cell = row.AddCell()
	cell.Value = strconv.Itoa(totalnc)
	row = sheet.AddRow()

	var personlist []string
	for k := range person {
		personlist = append(personlist, k)
	}
	sort.Strings(personlist)

	var communitylist []string
	for k := range communities {
		communitylist = append(communitylist, k)
	}
	sort.Strings(communitylist)

	for _, v := range communitylist {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = v
		for key, value := range personcommunity[v] {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = key
			cell = row.AddCell()
			cell.Value = strconv.Itoa(value["note"])
			cell = row.AddCell()
			cell.Value = strconv.Itoa(value["call"])
			cell = row.AddCell()
			var totalvalue int
			totalvalue = value["call"] + value["note"]
			cell.Value = strconv.Itoa(totalvalue)

		}
		row = sheet.AddRow()
	}

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Community"
	cell = row.AddCell()
	cell.Value = "Notes"
	cell = row.AddCell()
	cell.Value = "Calls"
	cell = row.AddCell()
	cell.Value = "Notes + Calls"

	for _, k := range communitylist {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = k
		cell = row.AddCell()
		cell.Value = strconv.Itoa(communitynotes[k])
		cell = row.AddCell()
		cell.Value = strconv.Itoa(communitycalls[k])
		cell = row.AddCell()
		cell.Value = strconv.Itoa(communities[k])
	}

	row = sheet.AddRow()
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Person"
	cell = row.AddCell()
	cell.Value = "Total Notes"
	cell = row.AddCell()
	cell.Value = "Total Calls"
	cell = row.AddCell()
	cell.Value = "Notes + Calls"

	for _, v := range personlist {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = v
		cell = row.AddCell()
		cell.Value = strconv.Itoa(personnotes[v])
		cell = row.AddCell()
		cell.Value = strconv.Itoa(personcalls[v])
		cell = row.AddCell()
		cell.Value = strconv.Itoa(person[v])
	}

	/*row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Person/Community"
	cell = row.AddCell()
	cell.Value = "Notes"
	cell = row.AddCell()
	cell.Value = "Calls"*/

	//for k, v := range personcommunity {

	//}


	err = file.Save("CombinedList" + time.Now().Local().Format("2006-01-02") + ".xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
