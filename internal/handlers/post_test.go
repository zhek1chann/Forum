package handlers

import (
	mock "forum/internal/repo/mocks"
	"net/http"
	"net/url"
	"testing"
)

func TestPost(t *testing.T) {
	ts := NewTestServer(t)

	defer ts.Close()

	tests := []struct {
		name     string
		url      string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			url:      "/post/1",
			wantCode: http.StatusOK,
		},
		{
			name:     "Negative ID",
			url:      "/post/-1",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Decimal ID",
			url:      "/post/1.77",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "String ID",
			url:      "/post/asdf",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Empty ID",
			url:      "/postid=view",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, _ := ts.get(t, tt.url)

			mock.Equal(t, code, tt.wantCode)
		})
	}
}

func TestCreatePost(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Close()

	const (
		validTitle   = "mihail"
		validContent = "petrovich"
	)
	validCategory := []string{"Python", "Javascript"}

	tests := []struct {
		name     string
		title    string
		content  string
		category []string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid post",
			title:    validTitle,
			content:  validContent,
			category: validCategory,
			wantCode: http.StatusSeeOther,
		},
		{
			name:     "Blank title",
			title:    "",
			content:  validContent,
			category: validCategory,
			wantCode: http.StatusSeeOther,
		},
		{
			name:     "Blank content",
			title:    validTitle,
			content:  "",
			category: validCategory,
			wantCode: http.StatusSeeOther,
		},
		{
			name:     "Too long title",
			title:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.",
			content:  validContent,
			category: validCategory,
			wantCode: http.StatusSeeOther,
		},
		{
			name:     "Too long content",
			title:    validTitle,
			content:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non libero placerat, ullamcorper purus in, hendrerit tellus. Etiam quis magna sagittis lorem tincidunt gravida. Suspendisse potenti. Cras tortor nisi, suscipit id ex eu, porta blandit arcu. Aliquam lacinia lorem est, sit amet tincidunt nisi fringilla non. Nullam sit amet quam nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed porta eget enim eu auctor. Cras ac maximus purus. Duis et tincidunt urna. Mauris nec quam sit amet massa tristique dapibus nec eu neque. Curabitur ut maximus lorem. Proin diam diam, ultricies ac condimentum nec, hendrerit at ex.",
			category: validCategory,
			wantCode: http.StatusSeeOther,
		},

		{
			name:     "No tags",
			title:    validTitle,
			content:  validContent,
			wantCode: http.StatusSeeOther,
		},
		{
			name:     "Invalid tags",
			title:    validTitle,
			content:  validContent,
			category: []string{"zaza"},
			wantCode: http.StatusSeeOther,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("title", tt.title)
			form.Add("content", tt.content)
			form["category"] = tt.category

			code, _, _ := ts.postForm(t, "/post/create", form)
			mock.Equal(t, code, tt.wantCode)
		})
	}
}
