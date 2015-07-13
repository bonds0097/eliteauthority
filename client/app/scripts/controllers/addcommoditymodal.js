'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:AddcommoditymodalCtrl
 * @description
 * # AddcommoditymodalCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
    .controller('AddCommodityModalCtrl', ['Commodities', 'gameDataService', '$modalInstance', '$scope', function(Commodities, gameDataService, $modalInstance, $scope) {
        var self = this;

        self.commodityCategories = gameDataService.commodityCategories();

        // Alert messages.
        self.alerts = [];
        self.closeAlert = function(index) {
            self.alerts.splice(index, 1);
        };

        self.reset = function(form) {
            if (form) {
                form.$setPristine();
                form.$setUntouched();
            }
            self.newCommodity = {};
        };

        self.cancel = function() {
            $modalInstance.dismiss('User canceled.');
        };

        self.saveCommodity = function(newCommodity) {
            var commodity = new Commodities(newCommodity);
            commodity.$save().then(function(response) {
                $modalInstance.close(response);
            }, function(errResponse) {
                self.alerts.push({
                    type: 'danger',
                    message: errResponse.data
                });
            });
        };

        self.reset();

        self.newCommodity.name = $scope.params.name;
        if ($scope.params.galacticAverage) {
            self.newCommodity.galacticAverage = $scope.params.galacticAverage;
        }
    }]);
