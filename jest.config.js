module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',
  moduleNameMapper: {
    '\\.(css|less|scss|sass)$': '<rootDir>/src/__mocks__/styleMock.js',
    '\\.(gif|ttf|eot|svg|png)$': '<rootDir>/src/__mocks__/fileMock.js'
  },
  setupFilesAfterEnv: ['<rootDir>/src/setupTests.ts'],
};