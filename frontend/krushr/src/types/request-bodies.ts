// Route
export interface PostRouteBody {
  name: string;
  pointsOfInterest: PostPointOfInterestBody[];
  imageIds?: number[];
  details?: PostDetailBody[];
  links?: PostLinkBody[];
  categories?: PostCategoryBody[];
  statusId: number;
}

export interface PostRouteBodyOld {
  name: string;
  pointOfInterestIds?: number[];
  imageIds?: number[];
  detailIds?: number[];
  linkIds?: number[];
  categoryIds?: number[];
  statusId: number;
}

export interface PutRouteBody extends PostRouteBody {}

// Image
export interface PostImageBody {
  path: string;
}

export interface PutImageBody extends PostImageBody {}

// Detail
export interface PostDetailBody {
  text: string;
}

export interface PutDetailBody extends PostDetailBody {}

// Link
export interface PostLinkBody {
  url: string;
}

export interface PutLinkBody extends PostLinkBody {}

// Category
export interface PostCategoryBody {
  name: string;
  position: number;
}

export interface PutCategoryBody extends PostCategoryBody {}

// Status
export interface PostStatusBody {
  name: string;
}

export interface PutStatusBody extends PostStatusBody {}

// PointOfInterest
export interface PostPointOfInterestBody extends GetPointOfInterestBody {
  imageIds?: number[];
  details?: PostDetailBody[];
  links?: PostLinkBody[];
  categories?: PostCategoryBody[];
}

export interface GetPointOfInterestBody {
  name: string;
  longitude: number;
  latitude: number;
}

export type PutPointOfInterestBody = PostPointOfInterestBody;

export interface PostUserBody {
  email: string;
  password: string;
  roleId: number;
}

export interface PutUserBody extends PostUserBody {}

export interface SignUpBody {
  email: string;
  password: string;
}

export interface SignInBody extends SignUpBody {}
