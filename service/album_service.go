package service

import (
	"errors"
	"gin_api/model"
)

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() []model.Album {
	return albums
}

func GetAlbumById(id string) (model.Album, error) {
	for _, album := range albums {
		if album.ID == id {
			return album, nil
		}
	}

	return model.Album{}, errors.New("NOT FOUND")
}

func PostAlbums(new_albums []model.Album) bool {
	albums = append(albums, new_albums...)

	return true
}
