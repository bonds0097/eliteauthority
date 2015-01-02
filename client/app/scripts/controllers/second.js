'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:SecondCtrl
 * @description
 * # SecondCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
  .controller('SecondCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
