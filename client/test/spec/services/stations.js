'use strict';

describe('Service: stations', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var stations;
  beforeEach(inject(function (_stations_) {
    stations = _stations_;
  }));

  it('should do something', function () {
    expect(!!stations).toBe(true);
  });

});
