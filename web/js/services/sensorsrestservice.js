app.service('sensorsRestService', [ '$resource',
    function($resource, globalParams) {
        return {
            action : function(params, callback, errorCallback) {
                $resource('/action', null, {
                    'get' : { method : 'POST'}
                }).get(params, callback, errorCallback);
            }
        };
    } ])