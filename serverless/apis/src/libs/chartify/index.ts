import {
  HEIGHT,
  MAX_VALUE,
  STROKE_COLOR,
  STROKE_WIDTH,
  STYLES,
  WIDTH,
} from './defaults';
import { ScaleOptions, SizeOptions, StrokeOptions } from './types';

export const chartify = (
  {
    data = [],
    scaleOptions: { minValue = 0, maxValue = MAX_VALUE } = {},
    sizeOptions: { width = WIDTH, height = HEIGHT } = {},
    stylesOptions = {},
    strokeOptions: {
      strokeColor = STROKE_COLOR,
      strokeWidth = STROKE_WIDTH,
    } = {},
  }: {
    data?: number[];
    scaleOptions?: ScaleOptions;
    sizeOptions?: SizeOptions;
    stylesOptions?: Record<string, string | number>;
    strokeOptions?: StrokeOptions;
  } = {
    data: [],
    scaleOptions: { minValue: 0, maxValue: MAX_VALUE },
    sizeOptions: { width: WIDTH, height: HEIGHT },
    stylesOptions: {},
    strokeOptions: {
      strokeColor: STROKE_COLOR,
      strokeWidth: STROKE_WIDTH,
    },
  }
): string => {
  const style: string = processStyle({
    sizeOptions: { width, height },
    stylesOptions,
  });
  const viewBox = `0 0 ${width} ${height}`;
  const points: string = processPoints({
    data,
    width,
    height,
    minValue,
    maxValue,
  });

  return `<svg viewBox="${viewBox}" style="${style}">
    <polyline
      fill="none"
      stroke="${strokeColor}"
      stroke-width="${strokeWidth}"
      points="${points}"
    />
</svg>`;
};

const processStyle = (
  {
    sizeOptions: { width = WIDTH, height = HEIGHT } = {},
    stylesOptions,
  }: {
    sizeOptions?: SizeOptions;
    stylesOptions?: Record<string, string | number>;
  } = {
    sizeOptions: { width: WIDTH, height: HEIGHT },
    stylesOptions: {},
  }
): string => {
  const fullStylesOptions: Record<string, string | number | undefined> = {
    ...STYLES,
    ...stylesOptions,
    width,
    height,
  };
  return Object.keys(fullStylesOptions)
    .map((key: string) => {
      const value = fullStylesOptions[key];
      if (typeof value === 'string') return `${key}: ${value}`;
      if (typeof value === 'number') return `${key}: ${value}px`;
      return null;
    })
    .filter((value) => value !== null)
    .join(';');
};

const processPoints = ({
  data = [],
  width = WIDTH,
  height = HEIGHT,
  minValue = 0,
  maxValue = MAX_VALUE,
}: {
  data: number[];
  width: number;
  height: number;
  minValue: number;
  maxValue: number;
}) => {
  const numberOfPoints = data.length;
  return data
    .map(
      (_value: number, index: number) => (width / (numberOfPoints - 1)) * index
    )
    .map((x: number, index: number) => {
      const value: number = data[index];
      const percentage = (value - minValue) / (maxValue - minValue);
      const y = height * (1 - percentage);
      return `${x},${y}`;
    })
    .join('\n');
};
