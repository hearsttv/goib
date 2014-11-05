package ib

import "testing"

func TestEntryApi(t *testing.T) {
	a := NewAPI("ibsys-c4k")

	httpClient := func(url string) ([]byte, error) {
		return []byte(entryJSON), nil
	}
	a.(*api).httpClient = httpClient

	coll, err := a.Entry("mostpopularcollection", nil)
	if err != nil {
		t.Errorf("error getting entry: %v", err)
	}
	if len(coll.Items) == 0 {
		t.Errorf("zero items returned from entry")
	}
}

func TestArticleApi(t *testing.T) {
	a := NewAPI("ibsys-c4k")

	httpClient := func(url string) ([]byte, error) {
		return []byte(articleJSON), nil
	}

	a.(*api).httpClient = httpClient

	article, err := a.Article(29341084, nil)
	if err != nil {
		t.Errorf("error getting content: %v", err)
	}
	if article.Type != ArticleType {
		t.Errorf("expected ARTICLE type but got %s", article.Type)
	}
	if article.ContentID != 29341084 {
		t.Errorf("expected 29341084 type but got %d", article.ContentID)
	}
	if article.Title != "6 things you should know about Diwali" {
		t.Errorf("expected '6 things you should know about Diwali' but got %s", article.Title)
	}
}

func TestVideoApi(t *testing.T) {
	a := NewAPI("ibsys-c4k")

	httpClient := func(url string) ([]byte, error) {
		return []byte(videoJSON), nil
	}

	a.(*api).httpClient = httpClient

	video, err := a.Video(1402356, nil)
	if err != nil {
		t.Errorf("error getting content: %v", err)
	}
	if video.Type != VideoType {
		t.Errorf("expected VIDEO type but got %s", video.Type)
	}
	if video.ContentID != 1402356 {
		t.Errorf("expected 1402356 type but got %d", video.ContentID)
	}
	if video.Title != "Advertisers Ready For NFL Kickoff" {
		t.Errorf("expected 'Advertisers Ready For NFL Kickoff' but got %s", video.Title)
	}
}

func TestSearchApi(t *testing.T) {
	a := NewAPI("ibsys-c4k")

	httpClient := func(url string) ([]byte, error) {
		return []byte(searchJSON), nil
	}
	a.(*api).httpClient = httpClient

	search, err := a.Search("nfl", nil)
	if err != nil {
		t.Errorf("error getting content: %v", err)
	}
	if search.TotalCount == 0 {
		t.Errorf("zero results returned from search")
	}
	if search.Keywords != "nfl" {
		t.Errorf("expected keywords 'nfl' but got %s", search.Keywords)
	}
}

var articleJSON = `
  {
  "type" : "ARTICLE",
  "hash" : "53v2bn",
  "content_id" : 29341084,
  "title" : "6 things you should know about Diwali",
  "subheadline" : "Hindu holiday celebrates good over evil",
  "teaser_title" : "6 things you should know about Diwali",
  "teaser_text" : "<p>The Hindu holiday of Diwali is India's biggest and brightest national holiday. But Indians around the world come together to celebrate the festival of lights.</p>",
  "teaser_image" : "http://www.channel4000.com/image/view/-/29345500/highRes/2/-/maxh/225/maxw/300/-/wxf17lz/-/Diwali-jpg.jpg",
  "article_text" : "<p>The Hindu holiday of Diwali is India's biggest and brightest national holiday. But Indians around the world come together to celebrate the festival of lights.</p><p>The five-day celebration of good over evil is as important to Hindus as Christmas is to Christians, and it marks the start of a new financial year for Indian businesses worldwide. But how much do you know about this global holiday, which began Thursday? Here are some facts and stats to help you improve your Diwali literacy.</p><p><strong>1. Diwali or Deepavali means rows of lights or lamps</strong></p><p>Diwali is known as the festival of lights because of the oil lamps and electric lights that people use to decorate homes, businesses and public spaces. As a celebration of the victory of good over evil and darkness over light, light is an important physical and spiritual symbol of the holiday.</p><p><strong>2. The name for a Hindu place of worship is \"mandir\"</strong></p><p>Christianity has churches, Judaism has synagogues, Islam has mosques and Hinduism has mandirs. On Diwali, Indians living abroad gather in mandirs for community celebrations. People leave offerings of food at the altars of different gods and gather for communal meals. Some mandirs host fireworks displays.</p><p><strong>3. Followers of various religions observe customs related to Diwali</strong></p><p>For many Indians, Diwali honors Lakshmi, the Hindu goddess of wealth. They light their homes and open their doors and windows to welcome her. In addition to Hindus, Jains, Buddhists and Sikhs also celebrate Diwali in such countries as Nepal, Bangladesh, Malaysia and Singapore. Legends and customs accompanying Diwali celebrations vary among religions and regions.</p><p><strong>4. People spend a lot of time getting ready</strong></p><p>Families spend days cleaning and decorate their homes in preparation for Diwali. They also shop for new clothes and outfits to wear to celebrations. Unsurprisingly, there are a lot of ways to go about this, from arts and crafts to makeup tricks to help you \"shimmer, shine and sparkle.\"</p><p><strong>5. Diwali is a big celebration in England, too</strong></p><p>This year, about 30,000 people attended a \"switch-on\" of more than 6,000 lamps in Leicester to mark the start of Diwali. Indians are the second-largest minority in Britain, according to IBT.</p><p><strong>6. An episode of \"The Office\" was dedicated to Diwali</strong></p><p>In one of few depictions of the holiday in American pop culture, bumbling Dunder Mifflin boss Michael Scott encourages his staff to support Kelly Kapoor by attending a local Diwali celebration. Cringe-worthy moments ensue.</p>",
  "media" : [ {
    "type" : "IMAGE",
    "hash" : "53xxbf",
    "content_id" : 29345500,
    "teaser_title" : "Diwali",
    "teaser_text" : null,
    "teaser_image" : "http://www.channel4000.com/image/view/-/29345500/highRes/2/-/maxh/225/maxw/300/-/wxf17lz/-/Diwali-jpg.jpg",
    "publication_date" : 1414323569
  } ],
  "related_media" : [ {
    "type" : "COLLECTION",
    "hash" : "101u3p7",
    "content_id" : 26117886,
    "teaser_title" : "",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1400770953
  } ],
  "keywords" : "Feature, Diwali",
  "categories" : [ {
    "title" : "",
    "hierarchy" : "/CNN Routine Priority",
    "path" : "/Shared Content/CNN Wire/Admin/CNN Routine Priority",
    "id" : "14228",
    "hash" : "t40f6j"
  }, {
    "title" : "",
    "hierarchy" : "/Master Parent/Lifestyle - Parent/National Family Headlines",
    "path" : "/Shared Content/IB News And Content/_EDITOR RESOURCES/IB Editorial Publishing Categories/Lifestyle/National Family Headlines",
    "id" : "13346",
    "hash" : "t3yqoj"
  }, {
    "title" : "",
    "hierarchy" : "/Master Parent/Lifestyle - Parent/National Lifestyle Headlines",
    "path" : "/Shared Content/IB News And Content/_EDITOR RESOURCES/IB Editorial Publishing Categories/Lifestyle/National Lifestyle Headlines",
    "id" : "71104",
    "hash" : "t7mjqk"
  }, {
    "title" : "",
    "hierarchy" : "/Master Parent/Lifestyle - Parent/National Travel Headlines",
    "path" : "/Shared Content/IB News And Content/_EDITOR RESOURCES/IB Editorial Publishing Categories/Lifestyle/National Travel Headlines",
    "id" : "13338",
    "hash" : "t3yqnp"
  } ],
  "copyright_objects" : [ {
    "name" : "2014-CNN Copyright",
    "text" : "<p><em>Copyright 2014 by <a href=\"http://www.cnn.com\">CNN NewSource</a>. All rights reserved. This material may not be published, broadcast, rewritten or redistributed.</em></p>"
  } ],
  "author_objects" : [ ],
  "author_location" : "(CNN)",
  "valid_from" : 2700000,
  "valid_to" : 1416873559,
  "creation_date" : 1414281560,
  "publication_date" : 1414316709,
  "editorial_comment" : "",
  "copyright" : "",
  "author" : "By Emanuella Grinberg CNN",
  "url" : "http://www.channel4000.com/travel/6-things-you-should-know-about-diwali/29341084",
  "navigation_context" : [ "IB News and Content -  Parent", "Shared Content", "CNN Wire", "Travel" ],
  "analytics_category" : "home",
  "advertising_category" : "homepage",
  "canonical_url" : "",
  "struct" : [ {
  } ]
}
`

var videoJSON = `
{
  "type" : "VIDEO",
  "hash" : "c4ym6r",
  "content_id" : 1402356,
  "title" : "Advertisers Ready For NFL Kickoff",
  "subheadline" : "",
  "teaser_title" : "Advertisers Ready For NFL Kickoff",
  "teaser_text" : "<p>Advertisers are gearing up for the NFL season as big name corporations shell out big bucks.\n</p>",
  "teaser_image" : null,
  "media" : [ ],
  "show_ads" : true,
  "caption" : null,
  "keywords" : "",
  "categories" : [ ],
  "copyright_objects" : [ ],
  "valid_from" : 2700000,
  "valid_to" : 4105144800,
  "creation_date" : 1315336037,
  "publication_date" : 1315336195,
  "editorial_comment" : null,
  "copyright" : "",
  "external_id" : "kaltura:0_063kr47w",
  "url" : "http://www.channel4000.com/Advertisers-Ready-For-NFL-Kickoff/1402356",
  "flavors" : [ {
    "video_type" : "flv",
    "url" : "http://kv.channel4000.com/p/557781/sp/55778100/download/entry_id/0_063kr47w/flavor/0_nqacuglv",
    "bitrate" : 456,
    "duration" : 78,
    "file_size" : 4925,
    "codec" : "vp6",
    "width" : 576,
    "height" : 324,
    "tags" : "web,source"
  } ],
  "m3u8" : "http://kv.channel4000.com/p/557781/sp/55778100/playManifest/entryId/0_063kr47w/format/applehttp/protocol/http/a.m3u8?deliveryCode=130835",
  "navigation_context" : [ "IB News and Content -  Parent", "Shared Content" ],
  "analytics_category" : "home",
  "advertising_category" : "homepage",
  "canonical_url" : "",
  "struct" : [ {
  } ]
}
`
var entryJSON = `
{
  "type" : "COLLECTION",
  "hash" : "1rhetx",
  "content_id" : 24208900,
  "teaser_title" : "Most Popular Stories",
  "teaser_image" : null,
  "collection_name" : "Most Popular Articles",
  "total_count" : 10,
  "start_index" : 0,
  "items" : [ {
    "type" : "ARTICLE",
    "hash" : "5o6xx5",
    "content_id" : 29461466,
    "teaser_title" : "1 pilot dead in Virgin's spaceship failure",
    "teaser_text" : "<p>The first sign there was a problem Friday with Virgin Galactic's SpaceShipTwo came at about 45,000 feet, just two minutes after the spaceplane separated from the jet-powered aircraft that carried it aloft, officials said. </p><p>It wasn't something overt wi...</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29467198/highRes/2/-/maxh/225/maxw/300/-/fr2cvoz/-/Map-of-SpacecraftTwo-crash-jpg.jpg",
    "publication_date" : 1414811176
  }, {
    "type" : "ARTICLE",
    "hash" : "5njs3q",
    "content_id" : 29450770,
    "teaser_title" : "Judge rejects Ebola quarantine for Maine nurse",
    "teaser_text" : "<p>A Maine judge on Friday ruled in favor of a nurse who defied a quarantine in a tense standoff with state authorities, saying local health officials failed to prove the need for a stricter order enforcing an Ebola quarantine. </p><p>District Court Chief Judg...</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29461910/highRes/2/-/maxh/225/maxw/300/-/wiithv/-/Kaci-Hickox--10-31-2014-jpg.jpg",
    "publication_date" : 1414795992
  }, {
    "type" : "ARTICLE",
    "hash" : "5njssp",
    "content_id" : 29450846,
    "teaser_title" : "Frein search ends, life resumes",
    "teaser_text" : "<p>Whew.</p><p>After 48 days of living with helicopters, heavily armed police officers, the rumors and constant questions -- where is he? -- the people of northeast Pennsylvania can finally relax.</p><p>The manhunt for Eric Matthew Frein, accused of killing a Pennsyl...</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29454620/highRes/2/-/maxh/225/maxw/300/-/itdjp1z/-/Eric-Frein-booking-photo-jpg.jpg",
    "publication_date" : 1414796929
  }, {
    "type" : "ARTICLE",
    "hash" : "5np9ii",
    "content_id" : 29458378,
    "teaser_title" : "Ax attack on D.C. cop prompts warnings",
    "teaser_text" : "<p>An officer with Washington's Metropolitan Police Department came under an unprovoked attack Friday from a man wielding an ax, police said.</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29461414/highRes/2/-/maxh/225/maxw/300/-/ls58h1/-/Ax-attack-on-cop-car-jpg.jpg",
    "publication_date" : 1414779219
  }, {
    "type" : "ARTICLE",
    "hash" : "5o7k2l",
    "content_id" : 29462042,
    "teaser_title" : "Chelsea Handler's topless photo fight",
    "teaser_text" : "<p>Comedian Chelsea Handler often aims to spark controversy, and one of her latest Instagram snapshots is no exception.</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29466856/highRes/2/-/maxh/225/maxw/300/-/n1x6saz/-/Chelsea-Handler-jpg.jpg",
    "publication_date" : 1414795750
  }, {
    "type" : "ARTICLE",
    "hash" : "5n3o2a",
    "content_id" : 29448178,
    "teaser_title" : "Sistine Chapel gets AC, lights makeover",
    "teaser_text" : "<p>High above the altar in the Vatican's Sistine Chapel, Michelangelo's masterpiece fresco is being seen in a new light..</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/18497606/highRes/2/-/maxh/225/maxw/300/-/52jdj0/-/Sistine-chapel-jpg.jpg",
    "publication_date" : 1414744143
  }, {
    "type" : "ARTICLE",
    "hash" : "5ort91",
    "content_id" : 29470548,
    "teaser_title" : "3rd victim dies from Washington school shooting",
    "teaser_text" : "<p>A third victim of last week's shooting at a Washington state high school has died, according to the hospital that treated victims.</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29326448/highRes/2/-/maxh/225/maxw/300/-/qbrmbu/-/Two-girls-after-Washington-school-shooting-jpg.jpg",
    "publication_date" : 1414808689
  }, {
    "type" : "ARTICLE",
    "hash" : "53v2bn",
    "content_id" : 29341084,
    "teaser_title" : "6 things you should know about Diwali",
    "teaser_text" : "<p>The Hindu holiday of Diwali is India's biggest and brightest national holiday. But Indians around the world come together to celebrate the festival of lights.</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29345500/highRes/2/-/maxh/225/maxw/300/-/wxf17lz/-/Diwali-jpg.jpg",
    "publication_date" : 1414316709
  }, {
    "type" : "ARTICLE",
    "hash" : "5mxald",
    "content_id" : 29440908,
    "teaser_title" : "NASCAR Driver Capsules, Oct. 30",
    "teaser_text" : "<p class=\"sdi-news-teaser\">Capsules for the eight drivers remaining in the Chase for the Sprint Cup.</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/85086/highRes/3/-/maxh/225/maxw/300/-/vn05e2/-/Racing-Generic.jpg",
    "publication_date" : 1414701807
  }, {
    "type" : "ARTICLE",
    "hash" : "5o8gkh",
    "content_id" : 29463998,
    "teaser_title" : "Cubs fire Renteria, open door to Maddon",
    "teaser_text" : "<p>The Chicago Cubs fired manager Rick Renteria on Friday, paving the way for the team to sign Joe Maddon.</p>",
    "teaser_image" : "http://www.channel4000.com/image/view/-/29322228/highRes/2/-/maxh/225/maxw/300/-/q5j0a9/-/Joe-Maddon--Tampa-Bay-Rays-manager-jpg.jpg",
    "publication_date" : 1414788360
  } ],
  "navigation_context" : [ "Home" ],
  "analytics_category" : "home",
  "advertising_category" : "homepage",
  "canonical_url" : "http://www.channel4000.com/24207306",
  "settings" : [ {
    "collection.limit" : "7"
  } ],
  "view_type" : "headlineStack",
  "struct" : [ {
  } ]
}
`

var searchJSON = `
{
  "type" : "SEARCH",
  "start_index" : 0,
  "total_count" : 4058,
  "keywords" : "nfl",
  "items" : [ {
    "type" : "COLLECTION",
    "hash" : "n3r8xa",
    "content_id" : 17901062,
    "teaser_title" : "NFL",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1356539088
  }, {
    "type" : "COLLECTION",
    "hash" : "u9g7hnz",
    "content_id" : 22566490,
    "teaser_title" : "NFL",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1382468080
  }, {
    "type" : "COLLECTION",
    "hash" : "ajqinnz",
    "content_id" : 15971136,
    "teaser_title" : "NFL",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1306891331
  }, {
    "type" : "COLLECTION",
    "hash" : "l1r7ut",
    "content_id" : 17571626,
    "teaser_title" : "NFL",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1313595682
  }, {
    "type" : "COLLECTION",
    "hash" : "rye1fyz",
    "content_id" : 19392694,
    "teaser_title" : "NFL",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1306891331
  }, {
    "type" : "COLLECTION",
    "hash" : "yi5ttg",
    "content_id" : 7929310,
    "teaser_title" : "",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1391443420
  }, {
    "type" : "COLLECTION",
    "hash" : "wyfo88z",
    "content_id" : 3979018,
    "teaser_title" : "NFL Headlines",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1319812007
  }, {
    "type" : "COLLECTION",
    "hash" : "1ynlik",
    "content_id" : 16182166,
    "teaser_title" : "NFL Headlines",
    "teaser_text" : null,
    "teaser_image" : null,
    "publication_date" : 1345295190
  }, {
    "type" : "IMAGE",
    "hash" : "t8emgg",
    "content_id" : 89020,
    "teaser_title" : "NFL logo",
    "teaser_text" : null,
    "teaser_image" : "http://www.channel4000.com/image/view/-/89020/highRes/3/-/maxh/225/maxw/300/-/rrhyhrz/-/NFL-Logo-jpg.jpg",
    "publication_date" : 1306260580
  }, {
    "type" : "IMAGE",
    "hash" : "30g59g",
    "content_id" : 132498,
    "teaser_title" : "NFL logo",
    "teaser_text" : null,
    "teaser_image" : "http://www.channel4000.com/image/view/-/132498/highRes/1/-/maxh/225/maxw/300/-/78n1rbz/-/NFL-Logo-jpg.jpg",
    "publication_date" : 1306260580
  } ]
}
`
