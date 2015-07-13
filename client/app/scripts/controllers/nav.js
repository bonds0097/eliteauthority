'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:NavctrlCtrl
 * @description
 * # NavctrlCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
    .controller('NavCtrl', ['$timeout', function($timeout) {
        var self = this;

        self.isCollapsed = false;

        self.showDropdown = function(dropdown) {
            $timeout.cancel(dropdown.timer);
            dropdown.timer = $timeout(function() {
                dropdown.isOpen = true;
            }, 50);
        };

        self.hideDropdown = function(dropdown) {
            $timeout.cancel(dropdown.timer);
            dropdown.timer = $timeout(function() {
                dropdown.isOpen = false;
            }, 500);
        };
    }]);
