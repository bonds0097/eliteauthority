'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:SecondCtrl
 * @description
 * # SecondCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
    .controller('SecondCtrl', ['SystemSearch', function(SystemSearch) {
        var self = this;

        self.findSystem = function(name) {
            var search = new SystemSearch({
                name: name
            });

            var matches = [];

            matches = search.$save().then(function(response) {
                if (response.systems === null) {
                    return [];
                }

                return response.systems;
            }, function(errResponse) {
                console.log(errResponse);
                return [];
            });

            return matches;
        };
    }]);
