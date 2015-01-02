'use strict';

describe('Service: Snapshots', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var Snapshots;
  beforeEach(inject(function (_Snapshots_) {
    Snapshots = _Snapshots_;
  }));

  it('should do something', function () {
    expect(!!Snapshots).toBe(true);
  });

});
