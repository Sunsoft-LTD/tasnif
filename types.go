package tasnif

const (
	API_URL = "https://tasnif.soliq.uz/api/cls-api"
)

type Language string

const (
	UZCYRL Language = "uz_cyrl"
	UZLATN Language = "uz_latn"
	RU     Language = "ru"
)

type (
	Tasnif struct {
		BaseURL  string
		Language Language
	}

	Product struct {
		MxikCode                    string `json:"mxikCode,omitempty"`
		Name                        string `json:"name,omitempty"`
		Description                 string `json:"description,omitempty"`
		InternationalCode           string `json:"internationalCode,omitempty"`
		Label                       string `json:"label,omitempty"`
		FullName                    string `json:"fullName,omitempty"`
		GroupCode                   string `json:"groupCode,omitempty"`
		GroupName                   string `json:"groupName,omitempty"`
		ClassCode                   string `json:"classCode,omitempty"`
		ClassName                   string `json:"className,omitempty"`
		PositionCode                string `json:"positionCode,omitempty"`
		PositionName                string `json:"positionName,omitempty"`
		SubPositionCode             string `json:"subPositionCode,omitempty"`
		SubPositionName             string `json:"subPositionName,omitempty"`
		BrandCode                   string `json:"brandCode,omitempty"`
		BrandName                   string `json:"brandName,omitempty"`
		AttributeName               string `json:"attributeName,omitempty"`
		UsePackage                  string `json:"usePackage,omitempty"`
		UnitsName                   string `json:"unitsName,omitempty"`
		SurveyCategoryID            string `json:"surveyCategoryId,omitempty"`
		NonChangeable               string `json:"nonChangeable,omitempty"`
		LgotaID                     string `json:"lgotaId,omitempty"`
		LgotaName                   string `json:"lgotaName,omitempty"`
		PackageName                 string `json:"packageName,omitempty"`
		UseCard                     string `json:"useCard,omitempty"`
		CategoryCode                string `json:"categoryCode,omitempty"`
		CategoryName                string `json:"categoryName,omitempty"`
		MnnName                     string `json:"mnnName,omitempty"`
		CategoryUnitID              any    `json:"categoryUnitId,omitempty"`
		CategoryUnitName            any    `json:"categoryUnitName,omitempty"`
		RecommendedCategoryUnitName any    `json:"recommendedCategoryUnitName,omitempty"`
		RecommendedUnitsName        any    `json:"recommendedUnitsName,omitempty"`
		Property                    any    `json:"property,omitempty"`
	}
	Response struct {
		Success     bool       `json:"success"`
		Code        int        `json:"code"`
		RecordTotal int        `json:"recordTotal"`
		Reason      string     `json:"reason"`
		Errors      any        `json:"error,omitempty"`
		Data        []*Product `json:"data"`
	}
)
