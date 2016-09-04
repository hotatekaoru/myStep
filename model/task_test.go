package model

import (
	"testing"
)

func TestConvOneTaskToS41(t *testing.T) {
	testConvTaskToS41Form1(t)
	testConvTaskToS41Form2(t)
	testConvTaskToS41Form3(t)
}

func testConvTaskToS41Form1(t *testing.T) {
	task := initSetS41Form1()
	form := convOneTaskToS41(*task)
	assertConvert1(t, &form)
}

func testConvTaskToS41Form2(t *testing.T) {
	task := initSetS41Form2()
	form := convOneTaskToS41(*task)
	assertConvert2(t, &form)
}

func testConvTaskToS41Form3(t *testing.T) {
	task := initSetS41Form3()
	form := convOneTaskToS41(*task)
	assertConvert3(t, &form)
}

func initSetS41Form1() *Task{
	task := Task{
		TypeId:      1,
		ContentId:   1,
		ContentName: "test",
		Point:       1.0,
		Par:         1,
		UnitId:      1,
	}
	return &task
}

func initSetS41Form2() *Task{
	task := Task{
		TypeId:      2,
		ContentId:   2,
		ContentName: "test",
		Point:       1.23,
		Par:         2,
		UnitId:      2,
	}
	return &task
}

func initSetS41Form3() *Task{
	task := Task{
		TypeId:      3,
		ContentId:   3,
		ContentName: "test",
		Point:       0.1,
		Par:         100,
		UnitId:      2,
	}
	return &task
}

func assertConvert1(t *testing.T, form *S41Form) {
	if "Coding" != form.TypeName {t.Errorf("assertConvert1.TypeName Error. TypeName is " + form.TypeName)}
	if "test" != form.ContentName {t.Errorf("assertConvert1.ContentName Error. TypeName is " + form.ContentName)}
	if "1pt / 1回" != form.PointStr {t.Errorf("assertConvert1.PointStr Error. TypeName is " + form.PointStr)}
}

func assertConvert2(t *testing.T, form *S41Form) {
	if "Training" != form.TypeName {t.Errorf("assertConvert1.TypeName Error. TypeName is " + form.TypeName)}
	if "1.23pt / 2分" != form.PointStr {t.Errorf("assertConvert1.PointStr Error. TypeName is " + form.PointStr)}
}

func assertConvert3(t *testing.T, form *S41Form) {
	if "Housework" != form.TypeName {t.Errorf("assertConvert1.TypeName Error. TypeName is " + form.TypeName)}
	if "0.1pt / 100分" != form.PointStr {t.Errorf("assertConvert1.PointStr Error. TypeName is " + form.PointStr)}
}