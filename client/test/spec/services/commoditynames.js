'use strict';

describe('Service: CommodityNames', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var CommodityNames;
  beforeEach(inject(function (_CommodityNames_) {
    CommodityNames = _CommodityNames_;
  }));

  it('should do something', function () {
    expect(!!CommodityNames).toBe(true);
  });

});
