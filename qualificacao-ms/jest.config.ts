/**
 * For a detailed explanation regarding each configuration property, visit:
 * https://jestjs.io/docs/configuration
 */

import type {Config} from 'jest';

const config: Config = {
  // Use ts-jest ESM preset so TypeScript files are transformed as ESM
  preset: 'ts-jest/presets/default-esm',
  testEnvironment: 'node',
  extensionsToTreatAsEsm: ['.ts'],
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { useESM: true }]
  },
  rootDir: '.',
  verbose: true,
  collectCoverage: true,
  coverageDirectory: 'coverage',
  coverageProvider: 'v8',
  moduleNameMapper: {
    // Allow resolving ESM .js extensions in imports emitted by ts-jest
    '^(\\.{1,2}\/.*)\\.js$': '$1',
    '@/(.*)': '<rootDir>/src/$1',
  },
  transformIgnorePatterns: [
    // Ignore most node_modules but allow @faker-js/faker and .pnpm to be transformed
    '<rootDir>/node_modules/(?!(@faker-js\\/faker|\\.pnpm))',
    '<rootDir>/build',
    '<rootDir>/coverage'
  ],
};

export default config;
