package service

import (
	"context"
	"errors"

	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/internal/model"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/pkg"
)

type LinkService struct {
	linkRepo *repository.LinkRepository
}

func NewLinkService(linkRepo *repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
	}
}

func (l *LinkService) CreateShortLink(ctx context.Context, userId int, link dto.CreateLinkReq) (dto.LinkDetailRes, error) {
	slug := link.Slug

	//kondisi kosong custom slug
	if slug == "" {
		for {
			randomSlug := pkg.GenRandomSlug()
			exists, err := l.linkRepo.IsSlugExist(ctx, randomSlug)
			if err != nil {
				return dto.LinkDetailRes{}, err
			}
			if !exists {
				slug = randomSlug
				break
			}
		}
	} else {
		// kondisi user isi slug
		isValid, errMsg := pkg.IsValidSlug(slug)
		if !isValid {
			return dto.LinkDetailRes{}, errors.New(errMsg)
		}
		//cek kondisi slug di database
		exists, err := l.linkRepo.IsSlugExist(ctx, slug)
		if err != nil {
			return dto.LinkDetailRes{}, err
		}
		if exists {
			return dto.LinkDetailRes{}, errors.New("Slug is Already taken by Another Link")
		}
	}
	//simpan data ke databse
	linkData, err := l.linkRepo.AddNewLink(ctx, model.Link{
		UserId:      userId,
		OriginalUrl: link.OriginalUrl,
		Slug:        slug,
	})
	if err != nil {
		return dto.LinkDetailRes{}, err
	}

	return dto.LinkDetailRes{
		Id:          linkData.Id,
		OriginalUrl: linkData.OriginalUrl,
		Slug:        linkData.Slug,
		ShortUrl:    "https://short.link/" + linkData.Slug,
		ClickCount:  linkData.ClickCount,
		CreatedAt:   *linkData.CreatedAt,
	}, nil
}
