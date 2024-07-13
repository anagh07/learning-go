package main

import "fmt"

type intmap map[string]int

func (m intmap) output () {
  fmt.Println(m)
}

func main () {
  companyUrls := map[string]string{
    "quran":  "quran.com",
    "google": "google.com",
  }
  fmt.Println(companyUrls)
  fmt.Println(companyUrls["quran"])

  companyUrls["linkedin"] = "linkedin.com"
  fmt.Println(companyUrls)
  delete(companyUrls, "linkedin")
  fmt.Println(companyUrls)

  ratings := make(map[string]float32, 2)
  ratings["The Boys"] = 7.1
  ratings["House of Dragons"] = 6.8
  ratings["The Bear"] = 8.1
  fmt.Println(ratings)

  contributions := make(intmap, 5)
  contributions["Anagh"] = 177
  contributions["Someone"] = 1233
  contributions.output()

  // Loop through a map
  for key, value := range contributions {
    fmt.Println("key: ", key)
    fmt.Println("value: ", value)
  }
}
