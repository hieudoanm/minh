import { useWindowSize } from '@hieudoanm/use-window-size';
import {
  geoMercator,
  geoPath,
  GeoProjection,
  json,
  select,
  selectAll,
} from 'd3';
import React, { useEffect, useState } from 'react';

enum Status {
  HERE = '#ef4444',
  PLANNED = '#3b82f6',
  VISITED = '#22c55e',
}

const visitedCountries: Record<string, string> = {
  australia: Status.VISITED,
  belgium: Status.PLANNED,
  brazil: Status.PLANNED,
  canada: Status.PLANNED,
  china: Status.VISITED,
  estonia: Status.VISITED,
  finland: Status.VISITED,
  france: Status.VISITED,
  germany: Status.VISITED,
  italy: Status.VISITED,
  japan: Status.PLANNED,
  netherlands: Status.VISITED,
  portugal: Status.PLANNED,
  thailand: Status.VISITED,
  'united-kingdom': Status.VISITED,
  'united-states': Status.PLANNED,
  vietnam: Status.HERE,
  singapore: Status.VISITED,
  'south-korea': Status.PLANNED,
  spain: Status.PLANNED,
  sweden: Status.VISITED,
};

const Map: React.FC<{ countries: any[] }> = ({ countries = [] }) => {
  const [data, setData] = useState<any>();
  const { width, height } = useWindowSize();

  useEffect(() => {
    const onLoad = async () => {
      // Load external data and boot
      const data: any = await json(
        'https://raw.githubusercontent.com/holtzy/D3-graph-gallery/master/DATA/world.geojson'
      );
      setData(data);
    };

    onLoad();
  }, []);

  useEffect(() => {
    if (!data) return;
    // The svg
    const svg = select('svg#map');
    const svgWidth: number =
      width || window.innerWidth || parseFloat(svg.attr('width'));
    const svgHeight: number =
      height || window.innerHeight || parseFloat(svg.attr('height'));
    selectAll('g').remove();

    // Map and projection
    const projection: GeoProjection = geoMercator()
      .center([0, 0]) // GPS of location to zoom on
      .scale(120) // This is like the zoom
      .translate([svgWidth / 2, svgHeight / 2]);

    // Draw the map
    const path: any = geoPath().projection(projection);
    svg
      .append('g')
      .selectAll('path')
      .data(data.features)
      .enter()
      .append('path')
      .attr('fill', (d: any) => {
        const country = countries.find((c) => c.cca3 === d.id);
        const common = country?.name?.common || '';
        const official = country?.name?.official || '';
        if (!common && !official) return '#ffffff';
        const key: string = common.toLowerCase().split(' ').join('-');
        const color = visitedCountries[key];
        return color || '#ffffff';
      })
      .attr('d', path)
      .attr('title', (d: any) => {
        const country = countries.find((c) => c.cca3 === d.id);
        return country?.name?.common || '';
      })
      .style('stroke', 'none');
  }, [width, height, data, countries]);

  return <svg id="map" style={{ width, height, overflow: 'hidden' }} />;
};

export default Map;
