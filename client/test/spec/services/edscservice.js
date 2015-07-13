'use strict';

describe('Service: edscService', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var edscService;
  beforeEach(inject(function (_edscService_) {
    edscService = _edscService_;
  }));

  it('should do something', function () {
    expect(!!edscService).toBe(true);
  });

});
