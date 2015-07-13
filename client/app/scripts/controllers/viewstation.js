'use strict';

angular.module('clientApp')
    .controller('ViewStationCtrl', ['$routeParams', 'lodash', 'Stations', function($routeParams, lodash, Stations) {
        var self = this;

        self.station = Stations.get({
            slug: $routeParams.slug
        }, function() {
            self.show = {
                commodities: lodash.contains(self.station.facilities, 'Commodities'),
                outfitting: lodash.contains(self.station.facilities, 'Outfitting'),
                shipyard: lodash.contains(self.station.facilities, 'Shipyard'),
            };

            self.commodityGrid = {};
            self.commodityGrid.data = self.station.snapshots;
            self.commodityGrid.columnDefs = [{
                name: 'Commodity',
                field: 'commodity.name'
            }, {
                name: 'Sell',
                field: 'sellPrice'
            }, {
                name: 'Buy',
                field: 'buyPrice'
            }, {
                name: 'Demand',
                field: 'demand'
            }, {
                name: 'Supply',
                field: 'supply'
            }, {
                name: 'Galactic Average',
                field: 'galacticAverage'
            }];
        });

        self.tabs = {
            basicData: {
                open: false,
                editing: true,
            }
        };
    }]);
