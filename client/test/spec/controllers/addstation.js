'use strict';

describe('Controller: AddstationCtrl', function () {

  // load the controller's module
  beforeEach(module('clientApp'));

  var AddstationCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    AddstationCtrl = $controller('AddstationCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
