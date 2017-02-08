app.controller('SensorsCtrl', ['$scope', 'sensorsRestService',
                             function($scope, sensorsRestService) {

	$scope.errors = [];

	function checkErrorAndShowMessage(error) {
        var isError = typeof error != "undefined" && error != null && error != "";

        if (isError) {
            alert("Server error" + (isError ? ": " + error : ""));
        }

        return isError;
    }

	function errorCallback(data) {
	    errorMessage();
    }

    $scope.statusDevices = {};

	$scope.m = {
	    r: 0,
	    g: 0,
        b: 0
    }

    function postCommand(command) {
        sensorsRestService.action(command, function(response) {

            if (checkErrorAndShowMessage(response.error)) {
                return;
            }

            var devices = [];

            angular.forEach(response.data, function(value, key) {

                if (value.DeviceType == "RgbLedDevice") {
                    $scope.m.r = value.R;
                    $scope.m.b = value.G;
                    $scope.m.b = value.B;
                }

                value.key = key;
                devices.push(value);
            });

            var priorityMap = {};
            priorityMap["DeviceManager"] = 1;
            priorityMap["RgbLedDevice"] = 2;
            priorityMap["OneStateDevice"] = 3;
            priorityMap["SystemPropery"] = 4;

            devices.sort(function(a, b) {
                var ap = priorityMap[a.DeviceType];
                var bp = priorityMap[b.DeviceType];

                if (typeof ap == "undefined") ap = 1000;
                if (typeof bp == "undefined") ap = 1000;

                return ap == bp ? 0 : ( ap > bp ? 1: -1);
            });

            $scope.statusDevices = devices;
        }, errorCallback )
    }

    $scope.saveRgb = function() {

        var command = {
            name: "change led status",
            action: "rgbLedState",
            r: $scope.m.r,
            g: $scope.m.g,
            b: $scope.m.b
        };

        postCommand(command);
    }

    $scope.refresh = function() {
        var command = {
            name: "request overall status",
            action: "status"
        };

        postCommand(command);
    }

    $scope.refresh();

}]);