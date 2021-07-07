/// <reference path="./custom.d.ts" />
// tslint:disable
/**
 * Audio Mastery API
 * Audio Mastery API definition
 *
 * OpenAPI spec version: 1.0.0
 *
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


import * as url from "url";
import * as portableFetch from "portable-fetch";
import { Configuration } from "./configuration";

const BASE_PATH = "http://localhost:3000".replace(/\/+$/, "");

/**
 *
 * @export
 */
export const COLLECTION_FORMATS = {
    csv: ",",
    ssv: " ",
    tsv: "\t",
    pipes: "|",
};

/**
 *
 * @export
 * @interface FetchAPI
 */
export interface FetchAPI {
    (url: string, init?: any): Promise<Response>;
}

/**
 *
 * @export
 * @interface FetchArgs
 */
export interface FetchArgs {
    url: string;
    options: any;
}

/**
 *
 * @export
 * @class BaseAPI
 */
export class BaseAPI {
    protected configuration: Configuration;

    constructor(configuration?: Configuration, protected basePath: string = BASE_PATH, protected fetch: FetchAPI = portableFetch) {
        if (configuration) {
            this.configuration = configuration;
            this.basePath = configuration.basePath || this.basePath;
        }
    }
};

/**
 *
 * @export
 * @class RequiredError
 * @extends {Error}
 */
export class RequiredError extends Error {
    name: "RequiredError"
    constructor(public field: string, msg?: string) {
        super(msg);
    }
}

/**
 *
 * @export
 * @interface Error
 */
export interface Error {
    /**
     *
     * @type {number}
     * @memberof Error
     */
    code?: number;
    /**
     *
     * @type {string}
     * @memberof Error
     */
    status?: string;
    /**
     *
     * @type {string}
     * @memberof Error
     */
    message?: string;
}

/**
 *
 * @export
 * @interface InlineResponse200
 */
export interface InlineResponse200 {
    /**
     *
     * @type {string}
     * @memberof InlineResponse200
     */
    title?: string;
    /**
     *
     * @type {string}
     * @memberof InlineResponse200
     */
    version?: string;
    /**
     *
     * @type {string}
     * @memberof InlineResponse200
     */
    description?: string;
}


/**
 * InfoApi - fetch parameter creator
 * @export
 */
export const InfoApiFetchParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get API info
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getInfo(options: any = {}): FetchArgs {
            const localVarPath = `/info`;
            const localVarUrlObj = url.parse(localVarPath, true);
            const localVarRequestOptions = Object.assign({ method: 'GET' }, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * InfoApi - functional programming interface
 * @export
 */
export const InfoApiFp = function(configuration?: Configuration) {
    return {
        /**
         * Get API info
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getInfo(options?: any): (fetch?: FetchAPI, basePath?: string) => Promise<InlineResponse200> {
            const localVarFetchArgs = InfoApiFetchParamCreator(configuration).getInfo(options);
            return (fetch: FetchAPI = portableFetch, basePath: string = BASE_PATH) => {
                return fetch(basePath + localVarFetchArgs.url, localVarFetchArgs.options).then((response) => {
                    if (response.status >= 200 && response.status < 300) {
                        return response.json();
                    } else {
                        throw response;
                    }
                });
            };
        },
    }
};

/**
 * InfoApi - factory interface
 * @export
 */
export const InfoApiFactory = function (configuration?: Configuration, fetch?: FetchAPI, basePath?: string) {
    return {
        /**
         * Get API info
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getInfo(options?: any) {
            return InfoApiFp(configuration).getInfo(options)(fetch, basePath);
        },
    };
};

/**
 * InfoApi - object-oriented interface
 * @export
 * @class InfoApi
 * @extends {BaseAPI}
 */
export class InfoApi extends BaseAPI {
    /**
     * Get API info
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof InfoApi
     */
    public getInfo(options?: any) {
        return InfoApiFp(this.configuration).getInfo(options)(this.fetch, this.basePath);
    }

}

