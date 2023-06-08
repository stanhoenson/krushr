import type { MarkerOptions } from "leaflet";

export interface ExtendedMarkerOptions extends MarkerOptions {
  position: number;
}

export type CoordinatesWithPosition = {
  latitude: number;
  longitude: number;
  position: number;
};
