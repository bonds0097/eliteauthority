'use strict';

describe('Service: CommoditiesService', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var CommoditiesService;
  beforeEach(inject(function (_CommoditiesService_) {
    CommoditiesService = _CommoditiesService_;
  }));

  it('should do something', function () {
    expect(!!CommoditiesService).toBe(true);
  });

});
