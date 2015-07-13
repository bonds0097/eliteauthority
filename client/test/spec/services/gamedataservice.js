'use strict';

describe('Service: gameDataService', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var gameDataService;
  beforeEach(inject(function (_gameDataService_) {
    gameDataService = _gameDataService_;
  }));

  it('should do something', function () {
    expect(!!gameDataService).toBe(true);
  });

});
