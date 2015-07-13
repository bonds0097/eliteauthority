'use strict';

describe('Controller: AddcommoditymodalCtrl', function () {

  // load the controller's module
  beforeEach(module('clientApp'));

  var AddcommoditymodalCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    AddcommoditymodalCtrl = $controller('AddcommoditymodalCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
