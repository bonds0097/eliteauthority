'use strict';

describe('Service: Systems', function () {

  // load the service's module
  beforeEach(module('clientApp'));

  // instantiate service
  var Systems;
  beforeEach(inject(function (_Systems_) {
    Systems = _Systems_;
  }));

  it('should do something', function () {
    expect(!!Systems).toBe(true);
  });

});
