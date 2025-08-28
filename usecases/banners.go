package usecases

import (
	"apis_service/domain"
	"context"
)

func (u *UseCases) GetProduct(
	ctx context.Context,
	input int32,
) (*domain.Product, error) {
	return u.repoProducts.Get(ctx, input, nil)
}

//func (u *UseCases) GetBanners(
//	ctx context.Context,
//	input *domain.BannersListRequest,
//) (*domain.BannersListResponse, error) {
//	banners, err := u.repoBanners.List(ctx, input, nil)
//	if err != nil {
//		return nil, errors.Wrap(err, "UseCases GetBanners List")
//	}
//
//	count, err := u.repoBanners.Count(ctx, input)
//
//	return &domain.BannersListResponse{
//		Banners: banners,
//		Count:   count,
//	}, errors.Wrap(err, "UseCases GetBanners Count")
//}

//func (u *UseCases) UpsertBanner(
//	ctx context.Context,
//	input *domain.BannerUpsertRequest,
//) (*domain.Banner, error) {
//	if input.ID == nil {
//		id, err := u.repoAutoIncrement.GetIncrement(ctx, domain.BannersCollection)
//		if err != nil {
//			return nil, errors.Wrap(err, "UseCases UpsertBanner GetIncrement")
//		}
//
//		input.ID = &id
//		input.GenDates()
//	} else {
//		input.Dates = domain.BannerDates{
//			Edited: time.Now().UTC(),
//		}
//	}
//
//	res, err := u.repoBanners.Upsert(ctx, input)
//	if err != nil {
//		return nil, errors.Wrap(err, "UseCases UpsertBanner Upsert")
//	}
//
//	_, err = u.cache.repoCustom.UpsertTime(ctx, domain.CacheVersion)
//	if err != nil {
//		return nil, errors.Wrap(err, "UseCase UpsertTime")
//	}
//
//	return res, nil
//}
//
//func (u *UseCases) DeleteBanner(
//	ctx context.Context,
//	input int32,
//) error {
//	return u.repoBanners.Delete(ctx, input)
//}
//
//func (u *UseCases) GetBannerPublic(
//	ctx context.Context,
//	input int32,
//) (*domain.BannerPublic, error) {
//	res, err := u.repoBanners.Get(ctx, input, &bson.M{"images": 1})
//	if err != nil {
//		return nil, errors.Wrap(err, "useCases GetBannerPublic Get")
//	}
//
//	if res != nil {
//		return &domain.BannerPublic{
//			ID:     res.ID,
//			Images: res.Images,
//		}, nil
//	}
//
//	return nil, nil
//}
//
//func (u *UseCases) GetBannersPublic(
//	ctx context.Context,
//	input *domain.BannersListRequest,
//) (*domain.BannersListPublicResponse, error) {
//	lang := slib.GetLanguageFromContext(ctx)
//
//	banners := u.cache.cacheBanners[input.BannerType]
//
//	if input.BannerType == domain.BannerTypePartners {
//		// avoid panic, shuffle elements, get first 5
//		if len(banners) > 0 {
//			rand.Shuffle(len(banners), func(i, j int) {
//				banners[i], banners[j] = banners[j], banners[i]
//			})
//		}
//
//		res := make([]*domain.BannerPublic, len(banners))
//
//		for i := 0; i < len(banners) && i < 4; i++ {
//			res = append(res, &domain.BannerPublic{ID: banners[i].ID, Images: banners[i].Images})
//		}
//
//		return &domain.BannersListPublicResponse{Banners: res}, nil
//	}
//
//	result := make([]*domain.BannerPublic, 0, len(banners))
//
//	for _, b := range banners {
//		if b == nil || !bannerMatchFilters(b, input) {
//			continue
//		}
//
//		result = append(result, &domain.BannerPublic{
//			ID:              b.ID,
//			Images:          bannerFilterImages(b.Images, lang),
//			BackgroundColor: b.BackgroundColor,
//			Placement:       b.Placement,
//		})
//	}
//
//	return &domain.BannersListPublicResponse{Banners: result}, nil
//}
//
//func bannerFilterImages(images []*domain.BannerImage, lang string) []*domain.BannerImage {
//	result := make([]*domain.BannerImage, 0, len(images))
//	for _, img := range images {
//		if img.Lang.String() == strings.ToUpper(lang) {
//			result = append(result, img)
//		}
//	}
//
//	return result
//}
//
//func bannerMatchFilters(banner *domain.Banner, request *domain.BannersListRequest) bool {
//	if banner == nil {
//		return false
//	}
//
//	if request == nil {
//		return true
//	}
//
//	return isStatusMatch(banner.Status, request.Status) &&
//		isProjectMatch(banner.Project, request.BannerProject) &&
//		isPlatformMatch(banner.Platforms, request.Platform) &&
//		isAudienceMatch(banner.Audiences, request.Audience)
//}
//
//func isStatusMatch(bannerStatus, requestStatus domain.BannerStatus) bool {
//	return requestStatus == domain.BannerStatusUnspecified || requestStatus == bannerStatus
//}
//
//func isProjectMatch(bannerProject, requestProject domain.BannerProject) bool {
//	return requestProject == domain.BannerProjectUnspecified || requestProject == bannerProject
//}
//
//func isPlatformMatch(bannerPlatforms []domain.BannerPlatform, requestPlatform domain.BannerPlatform) bool {
//	return requestPlatform == domain.BannerPlatformUnspecified || slices.Contains(bannerPlatforms, requestPlatform)
//}
//
//func isAudienceMatch(bannerAudiences []domain.BannerAudience, requestAudience domain.BannerAudience) bool {
//	if slices.Contains(bannerAudiences, domain.BannerAudienceAll) {
//		return true
//	}
//
//	return requestAudience == domain.BannerAudienceUnspecified || slices.Contains(bannerAudiences, requestAudience)
//}
