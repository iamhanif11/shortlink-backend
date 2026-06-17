package service

import (
	"context"
	"errors"
	"time"

	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/internal/model"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/pkg"
	"github.com/redis/go-redis/v9"
)

type LinkService struct {
	linkRepo *repository.LinkRepository
	rc       *redis.Client
}

func NewLinkService(linkRepo *repository.LinkRepository, rc *redis.Client) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
		rc:       rc,
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

func (l *LinkService) GetUserLinks(ctx context.Context, userId int) ([]dto.LinkDetailRes, error) {
	linksData, err := l.linkRepo.GetLinksByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	//inisialisasi
	var result []dto.LinkDetailRes = []dto.LinkDetailRes{}

	for _, link := range linksData {
		var createdAtTime time.Time
		if link.CreatedAt != nil {
			createdAtTime = *link.CreatedAt
		} else {
			createdAtTime = time.Now()
		}

		dtoLink := dto.LinkDetailRes{
			Id:          link.Id,
			OriginalUrl: link.OriginalUrl,
			Slug:        link.Slug,
			ShortUrl:    "https://short.link" + link.Slug,
			ClickCount:  link.ClickCount,
			CreatedAt:   createdAtTime,
		}
		result = append(result, dtoLink)
	}
	return result, nil
}

func (l *LinkService) DeleteLink(ctx context.Context, id, userId int) error {
	slug, err := l.linkRepo.SoftDeleteLink(ctx, id, userId)
	if err != nil {
		return err
	}

	//cache invalidation
	redisKey := "link: " + slug
	_ = l.rc.Del(ctx, redisKey).Err()

	return nil
}

func (l *LinkService) RedirectSlug(ctx context.Context, slug string) (string, error) {
	originalUrl, err := l.linkRepo.GetAndIncrementClick(ctx, slug)
	if err != nil {
		return "", nil
	}
	return originalUrl, nil
}
