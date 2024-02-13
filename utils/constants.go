package utils

type ReadOnlyStringMap struct {
	data map[string]string
}

var DataciteToCommonMeta ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Audiovisual":           "Audiovisual",
		"BlogPosting":           "Article",
		"Book":                  "Book",
		"BookChapter":           "BookChapter",
		"Collection":            "Collection",
		"ComputationalNotebook": "ComputationalNotebook",
		"ConferencePaper":       "ProceedingsArticle",
		"ConferenceProceeding":  "Proceedings",
		"DataPaper":             "JournalArticle",
		"Dataset":               "Dataset",
		"Dissertation":          "Dissertation",
		"Event":                 "Event",
		"Image":                 "Image",
		"Instrument":            "Instrument",
		"InteractiveResource":   "InteractiveResource",
		"Journal":               "Journal",
		"JournalArticle":        "JournalArticle",
		"Model":                 "Model",
		"OutputManagementPlan":  "OutputManagementPlan",
		"PeerReview":            "PeerReview",
		"PhysicalObject":        "PhysicalObject",
		"Poster":                "Speech",
		"Preprint":              "Article",
		"Report":                "Report",
		"Service":               "Service",
		"Software":              "Software",
		"Sound":                 "Sound",
		"Standard":              "Standard",
		"Text":                  "Document",
		"Thesis":                "Dissertation",
		"Workflow":              "Workflow",
		"Other":                 "Other",
	},
}

var ResourceTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Article":            "Article",
		"Audiovisual":        "Audiovisual",
		"Book":               "Book",
		"BookChapter":        "BookChapter",
		"BookSeries":         "BookSeries",
		"Component":          "Component",
		"Dataset":            "Dataset",
		"Dissertation":       "Dissertation",
		"Document":           "Document",
		"Journal":            "Journal",
		"ProceedingsSeries":  "ProceedingsSeries",
		"Grant":              "Grant",
		"Instrument":         "Instrument",
		"JournalArticle":     "JournalArticle",
		"JournalIssue":       "JournalIssue",
		"JournalVolume":      "JournalVolume",
		"PeerReview":         "PeerReview",
		"PhysicalObject":     "PhysicalObject",
		"Proceedings":        "Proceedings",
		"ProceedingsArticle": "ProceedingsArticle",
		"Report":             "Report",
		"ReportComponent":    "ReportComponent",
		"ReportSeries":       "ReportSeries",
		"Software":           "Software",
		"StudyRegistration":  "StudyRegistration",
		"Other":              "Other",
	},
}

var RoleMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Author":                "Author",
		"Editor":                "Editor",
		"Chair":                 "Chair",
		"Reviewer":              "Reviewer",
		"ReviewAssistant":       "ReviewAssistant",
		"StatsReviewer":         "StatsReviewer",
		"ReviewerExternal":      "ReviewerExternal",
		"Reader":                "Reader",
		"Translator":            "Translator",
		"ContactPerson":         "ContactPerson",
		"DataCollector":         "DataCollector",
		"DataManager":           "DataManager",
		"Distributor":           "Distributor",
		"HostingInstitution":    "HostingInstitution",
		"Producer":              "Producer",
		"ProjectLeader":         "ProjectLeader",
		"ProjectManager":        "ProjectManager",
		"ProjectMember":         "ProjectMember",
		"RegistrationAgency":    "RegistrationAgency",
		"RegistrationAuthority": "RegistrationAuthority",
		"RelatedPerson":         "RelatedPerson",
		"ResearchGroup":         "ResearchGroup",
		"RightsHolder":          "RightsHolder",
		"Researcher":            "Researcher",
		"Sponsor":               "Sponsor",
		"WorkPackageLeader":     "WorkPackageLeader",
		"Conceptualization":     "Conceptualization",
		"DataCuration":          "DataCuration",
		"FormalAnalysis":        "FormalAnalysis",
		"FundingAcquisition":    "FundingAcquisition",
		"Investigation":         "Investigation",
		"Methodology":           "Methodology",
		"ProjectAdministration": "ProjectAdministration",
		"Resources":             "Resources",
		"Software":              "Software",
		"Supervision":           "Supervision",
		"Validation":            "Validation",
		"Visualization":         "Visualization",
		"WritingOriginalDraft":  "WritingOriginalDraft",
		"WritingReviewEditing":  "WritingReviewEditing",
		"Maintainer":            "Maintainer",
		"Other":                 "Other",
	},
}

var ArchiveLocationMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Clockss":         "CLOCKSS",
		"Dwt":             "DWT",
		"InternetArchive": "Internet Archive",
		"KB":              "KB",
		"Lockss":          "LOCKSS",
		"Portico":         "Portico",
	},
}

var ContainerTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"DataCatalog":             "DataCatalog",
		"Periodical":              "Periodical",
		"PurpleBook":              "Book",
		"PurpleBookSeries":        "BookSeries",
		"PurpleJournal":           "Journal",
		"PurpleProceedingsSeries": "ProceedingsSeries",
		"Repository":              "Repository",
		"Series":                  "Series",
	},
}

var ContributorTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Organization": "Organization",
		"Person":       "Person",
	},
}

var DescriptionTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Abstract":      "Abstract",
		"Methods":       "Methods",
		"Summary":       "Summary",
		"TechnicalInfo": "TechnicalInfo",
		"Other":         "Other",
	},
}

var FunderIdentifierTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"CrossrefFunderID":          "Crossref Funder ID",
		"FunderIdentifierTypeOther": "Other",
		"Grid":                      "GRID",
		"Isni":                      "ISNI",
		"Ringgold":                  "Ringgold",
		"Ror":                       "ROR",
	},
}

var ProviderTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Crossref":   "Crossref",
		"DataCite":   "DataCite",
		"GitHub":     "GitHub",
		"JaLC":       "JaLC",
		"Kisti":      "KISTI",
		"MEDRA":      "mEDRA",
		"Op":         "OP",
		"InvenioRDM": "InvenioRDM",
	},
}

var RelatedIdentifierTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"HasPart":             "HasPart",
		"HasPreprint":         "HasPreprint",
		"HasVersion":          "HasVersion",
		"IsIdenticalTo":       "IsIdenticalTo",
		"IsNewVersionOf":      "IsNewVersionOf",
		"IsOriginalFormOf":    "IsOriginalFormOf",
		"IsPartOf":            "IsPartOf",
		"IsPreprintOf":        "IsPreprintOf",
		"IsPreviousVersionOf": "IsPreviousVersionOf",
		"IsReviewedBy":        "IsReviewedBy",
		"IsSupplementTo":      "IsSupplementTo",
		"IsTranslationOf":     "IsTranslationOf",
		"IsVariantFormOf":     "IsVariantFormOf",
		"IsVersionOf":         "IsVersionOf",
		"Reviews":             "Reviews",
	},
}

var SchemaVersionMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"HTTPDataciteOrgSchemaKernel3":     "http://datacite.org/schema/kernel-3",
		"HTTPDataciteOrgSchemaKernel4":     "http://datacite.org/schema/kernel-4",
		"HTTPSCommonmetaOrgCommonmetaV011": "https://commonmeta.org/commonmeta_v0.11",
	},
}

var StateMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"Findable": "findable",
		"NotFound": "not_found",
	},
}

var TitleTypeMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"AlternativeTitle": "AlternativeTitle",
		"Subtitle":         "Subtitle",
		"TranslatedTitle":  "TranslatedTitle",
	},
}

var NormalizedCCLicensesMap ReadOnlyStringMap = ReadOnlyStringMap{
	data: map[string]string{
		"https://creativecommons.org/licenses/by/1.0":                 "https://creativecommons.org/licenses/by/1.0/legalcode",
		"https://creativecommons.org/licenses/by/2.0":                 "https://creativecommons.org/licenses/by/2.0/legalcode",
		"https://creativecommons.org/licenses/by/2.5":                 "https://creativecommons.org/licenses/by/2.5/legalcode",
		"https://creativecommons.org/licenses/by/3.0":                 "https://creativecommons.org/licenses/by/3.0/legalcode",
		"https://creativecommons.org/licenses/by/3.0/us":              "https://creativecommons.org/licenses/by/3.0/legalcode",
		"https://creativecommons.org/licenses/by/4.0":                 "https://creativecommons.org/licenses/by/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/1.0":              "https://creativecommons.org/licenses/by-nc/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/2.0":              "https://creativecommons.org/licenses/by-nc/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/2.5":              "https://creativecommons.org/licenses/by-nc/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nc/3.0":              "https://creativecommons.org/licenses/by-nc/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/4.0":              "https://creativecommons.org/licenses/by-nc/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/1.0":           "https://creativecommons.org/licenses/by-nd-nc/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/2.0":           "https://creativecommons.org/licenses/by-nd-nc/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/2.5":           "https://creativecommons.org/licenses/by-nd-nc/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/3.0":           "https://creativecommons.org/licenses/by-nd-nc/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/4.0":           "https://creativecommons.org/licenses/by-nd-nc/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/1.0":           "https://creativecommons.org/licenses/by-nc-sa/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/2.0":           "https://creativecommons.org/licenses/by-nc-sa/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/2.5":           "https://creativecommons.org/licenses/by-nc-sa/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/3.0":           "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/3.0/us":        "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/4.0":           "https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/1.0":              "https://creativecommons.org/licenses/by-nd/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/2.0":              "https://creativecommons.org/licenses/by-nd/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/2.5":              "https://creativecommons.org/licenses/by-nd/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nd/3.0":              "https://creativecommons.org/licenses/by-nd/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/4.0":              "https://creativecommons.org/licenses/by-nd/2.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/1.0":              "https://creativecommons.org/licenses/by-sa/1.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/2.0":              "https://creativecommons.org/licenses/by-sa/2.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/2.5":              "https://creativecommons.org/licenses/by-sa/2.5/legalcode",
		"https://creativecommons.org/licenses/by-sa/3.0":              "https://creativecommons.org/licenses/by-sa/3.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/4.0":              "https://creativecommons.org/licenses/by-sa/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/1.0":           "https://creativecommons.org/licenses/by-nc-nd/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/2.0":           "https://creativecommons.org/licenses/by-nc-nd/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/2.5":           "https://creativecommons.org/licenses/by-nc-nd/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/3.0":           "https://creativecommons.org/licenses/by-nc-nd/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/4.0":           "https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode",
		"https://creativecommons.org/licenses/publicdomain":           "https://creativecommons.org/licenses/publicdomain/",
		"https://creativecommons.org/publicdomain/zero/1.0":           "https://creativecommons.org/publicdomain/zero/1.0/legalcode",
		"https://creativecommons.org/licenses/by/1.0/legalcode":       "https://creativecommons.org/licenses/by/1.0/legalcode",
		"https://creativecommons.org/licenses/by/2.0/legalcode":       "https://creativecommons.org/licenses/by/2.0/legalcode",
		"https://creativecommons.org/licenses/by/2.5/legalcode":       "https://creativecommons.org/licenses/by/2.5/legalcode",
		"https://creativecommons.org/licenses/by/3.0/legalcode":       "https://creativecommons.org/licenses/by/3.0/legalcode",
		"https://creativecommons.org/licenses/by/4.0/legalcode":       "https://creativecommons.org/licenses/by/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/1.0/legalcode":    "https://creativecommons.org/licenses/by-nc/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/2.0/legalcode":    "https://creativecommons.org/licenses/by-nc/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/2.5/legalcode":    "https://creativecommons.org/licenses/by-nc/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nc/3.0/legalcode":    "https://creativecommons.org/licenses/by-nc/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc/4.0/legalcode":    "https://creativecommons.org/licenses/by-nc/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/1.0/legalcode": "https://creativecommons.org/licenses/by-nd-nc/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/2.0/legalcode": "https://creativecommons.org/licenses/by-nd-nc/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/2.5/legalcode": "https://creativecommons.org/licenses/by-nd-nc/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/3.0/legalcode": "https://creativecommons.org/licenses/by-nd-nc/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nd-nc/4.0/legalcode": "https://creativecommons.org/licenses/by-nd-nc/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/1.0/legalcode": "https://creativecommons.org/licenses/by-nc-sa/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/2.0/legalcode": "https://creativecommons.org/licenses/by-nc-sa/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/2.5/legalcode": "https://creativecommons.org/licenses/by-nc-sa/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode": "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode": "https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/1.0/legalcode":    "https://creativecommons.org/licenses/by-nd/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/2.0/legalcode":    "https://creativecommons.org/licenses/by-nd/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/2.5/legalcode":    "https://creativecommons.org/licenses/by-nd/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nd/3.0/legalcode":    "https://creativecommons.org/licenses/by-nd/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nd/4.0/legalcode":    "https://creativecommons.org/licenses/by-nd/2.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/1.0/legalcode":    "https://creativecommons.org/licenses/by-sa/1.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/2.0/legalcode":    "https://creativecommons.org/licenses/by-sa/2.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/2.5/legalcode":    "https://creativecommons.org/licenses/by-sa/2.5/legalcode",
		"https://creativecommons.org/licenses/by-sa/3.0/legalcode":    "https://creativecommons.org/licenses/by-sa/3.0/legalcode",
		"https://creativecommons.org/licenses/by-sa/4.0/legalcode":    "https://creativecommons.org/licenses/by-sa/4.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/1.0/legalcode": "https://creativecommons.org/licenses/by-nc-nd/1.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/2.0/legalcode": "https://creativecommons.org/licenses/by-nc-nd/2.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/2.5/legalcode": "https://creativecommons.org/licenses/by-nc-nd/2.5/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/3.0/legalcode": "https://creativecommons.org/licenses/by-nc-nd/3.0/legalcode",
		"https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode": "https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode",
		"https://creativecommons.org/licenses/publicdomain/":          "https://creativecommons.org/licenses/publicdomain/",
		"https://creativecommons.org/publicdomain/zero/1.0/legalcode": "https://creativecommons.org/publicdomain/zero/1.0/legalcode",
	},
}

// Get returns the matching value for given key. If the key is not in map, it
// may return "Other" or "".
func (r ReadOnlyStringMap) GetVal(k string, tryOther bool) string {
	if _, ok := r.data[k]; ok {
		return r.data[k]
	}
	if tryOther {
		if _, ok := r.data["Other"]; ok {
			return r.data["Other"]
		}
	}
	return ""
}