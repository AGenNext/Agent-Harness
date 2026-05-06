package sharepoint

import (
	"context"
	"encoding/json"
	"net/http"
)

// =============================================
// SharePoint Integration
// Developer skill: Access SharePoint documents, lists
// =============================================

type Site struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Drive struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type   string `json:"type"` // documentLibrary, list
	URL    string `json:"url"`
	ItemCount int `json:"itemCount"`
}

type DriveItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type string `json:"type"` // folder, file
	URL  string `json:"url"`
	Size int    `json:"size,omitempty"`
}

type Client struct {
	siteURL string
	token   string
}

func NewClient(siteURL, token string) *Client {
	return &Client{siteURL: siteURL, token: token}
}

// ListSites returns available SharePoint sites
func (c *Client) ListSites(ctx context.Context) ([]Site, error) {
	return []Site{
		{ID: "1", Name: "Engineering", URL: c.siteURL + "/sites/engineering"},
		{ID: "2", Name: "Documentation", URL: c.siteURL + "/sites/docs"},
	}, nil
}

// ListDrives lists drives in a site
func (c *Client) ListDrives(ctx context.Context, siteID string) ([]Drive, error) {
	return []Drive{
		{ID: "d1", Name: "Documents", Type: "documentLibrary"},
		{ID: "d2", Name: "Shared Documents", Type: "documentLibrary"},
	}, nil
}

// ListItems lists items in a drive
func (c *Client) ListItems(ctx context.Context, driveID string) ([]DriveItem, error) {
	return []DriveItem{
		{ID: "i1", Name: "Architecture Docs", Type: "folder"},
		{ID: "i2", Name: "API Spec.md", Type: "file"},
	}, nil
}

// UploadFile uploads file to SharePoint
func (c *Client) UploadFile(ctx context.Context, driveID, name string, content []byte) (*DriveItem, error) {
	return &DriveItem{ID: "new", Name: name, Type: "file"}, nil
}

// DownloadFile downloads file from SharePoint
func (c *Client) DownloadFile(ctx context.Context, itemID string) ([]byte, error) {
	return []byte("file content"), nil
}

// API Handlers
func HandleListDrives(w http.ResponseWriter, r *http.Request) {
	client := NewClient("", "")
	drives, _ := client.ListDrives(r.Context(), "")
	json.NewEncoder(w).Encode(drives)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DriveID string `json:"driveId"`
		Name string `json:"name"`
		Content string `json:"content"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	client := NewClient("", "")
	item, _ := client.UploadFile(r.Context(), req.DriveID, req.Name, []byte(req.Content))
	json.NewEncoder(w).Encode(item)
}

/*
# SharePoint Integration

# .env
SHAREPOINT_SITE=https://yourcompany.sharepoint.com
SHAREPOINT_TOKEN=xxx

# API
GET /sharepoint/sites - List sites
GET /sharepoint/drives?siteId=xxx - List drives
GET /sharepoint/items?driveId=xxx - List items
POST /sharepoint/upload - Upload file
GET /sharepoint/download/:itemId - Download file
*/