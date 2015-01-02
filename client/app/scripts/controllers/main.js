'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
    .controller('MainCtrl', ['Commodities', 'Snapshots', function(Commodities, Snapshots) {
        var self = this;

        self.dummyCommodity = function() {
            var dummy = new Commodities({
                name: 'Test Commodity ' + Math.floor(Math.random() * 10),
                category: 'Test Category ' + Math.floor(Math.random() * 25),
                galacticAverage: Math.floor(Math.random() * 100000)
            });

            dummy.$save().then(function(response) {
                console.log(response);
            }, function(errResponse) {
                console.log(errResponse);
            });
        };

        self.getCommodities = function() {
            Commodities.query().$promise.then(function(response) {
                for (var i = 0; i < response.length; i++) {
                    console.log(response[i].name);
                }
            }, function(errResponse) {
                console.log(errResponse);
            });
        };

        self.dummySnapshot = function() {
            var dummy = new Snapshots({
                commodity: 'Test Commodity ' + Math.floor(Math.random() * 10),
                station: 'Test Stations ' + Math.floor(Math.random() * 10),
                sellPrice: Math.floor(Math.random() * 100000),
                buyPrice: Math.floor(Math.random() * 100000),
                demand: Math.floor(Math.random() * 100000),
                supply: Math.floor(Math.random() * 100000),
                galacticAverage: Math.floor(Math.random() * 100000)
            });

            dummy.$save().then(function(response) {
                console.log(response);
            }, function(errResponse) {
                console.log(errResponse);
            });
        };
    }]);
