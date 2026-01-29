package service

import "github.com/FatManlife/component-finder/back-end/internal/models/dto"

func checkDefaultParams(prams *dto.ProductParams) {
	//Cheking min and max values
	if prams.Min > prams.Max {
		prams.Min = 0
		prams.Max = 0
	}

	if prams.Min < 0 {
		prams.Min = 0
	}

	if prams.Max < 0 {
		prams.Max = 0
	}

	// limit
	if prams.Limit <= 0 {
		prams.Limit = 20
	}

	// after
	if prams.After < 0 {
		prams.After = 0
	}
}