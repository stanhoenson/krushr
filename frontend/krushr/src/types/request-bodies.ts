// Route
interface PostRouteBody {
  name: string;
  pointsOfInterest: PostPointOfInterestBody[];
  imageIds?: number[];
  details?: PostDetailBody[];
  links?: PostLinkBody[];
  categories?: PostCategoryBody[];
  statusId: number;
}

interface PostRouteBodyOld {
  name: string;
  pointOfInterestIds?: number[];
  imageIds?: number[];
  detailIds?: number[];
  linkIds?: number[];
  categoryIds?: number[];
  statusId: number;
}

interface PutRouteBody extends PostRouteBody {
  id: number;
}

// Image
interface PostImageBody {
  path: string;
}

interface PutImageBody extends PostImageBody {
  id: number;
}

// Detail
interface PostDetailBody {
  text: string;
}

interface PutDetailBody extends PostDetailBody {
  id: number;
}

// Link
interface PostLinkBody {
  url: string;
}

interface PutLinkBody extends PostLinkBody {
  id: number;
}

// Category
interface PostCategoryBody {
  name: string;
  position: number;
}

interface PutCategoryBody extends PostCategoryBody {
  id: number;
}

// Status
interface PostStatusBody {
  name: string;
}

interface PutStatusBody extends PostStatusBody {
  id: number;
}

// PointOfInterest
interface PostPointOfInterestBody {
  name: string;
  longitude: number;
  latitude: number;
  imageIds?: number[];
  details?: PostDetailBody[];
  links?: PostLinkBody[];
  categories?: PostCategoryBody[];
}

interface PutPointOfInterestBody extends PostPointOfInterestBody {
  id: number;
}
