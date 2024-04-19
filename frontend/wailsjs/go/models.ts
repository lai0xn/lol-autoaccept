export namespace lcu {
	
	export class SunnmonerInfo {
	    accountId: number;
	    displayName: string;
	    gameName: string;
	    internalName: string;
	    nameChangeFlag: boolean;
	    percentCompleteForNextLevel: number;
	    privacy: string;
	    profileIconId: number;
	    puuid: string;
	    // Go type: struct { CurrentPoints int "json:\"currentPoints\""; MaxRolls int "json:\"maxRolls\""; NumberOfRolls int "json:\"numberOfRolls\""; PointsCostToRoll int "json:\"pointsCostToRoll\""; PointsToReroll int "json:\"pointsToReroll\"" }
	    rerollPoints: any;
	    summonerId: number;
	    summonerLevel: number;
	    tagLine: string;
	    unnamed: boolean;
	    xpSinceLastLevel: number;
	    xpUntilNextLevel: number;
	
	    static createFrom(source: any = {}) {
	        return new SunnmonerInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.displayName = source["displayName"];
	        this.gameName = source["gameName"];
	        this.internalName = source["internalName"];
	        this.nameChangeFlag = source["nameChangeFlag"];
	        this.percentCompleteForNextLevel = source["percentCompleteForNextLevel"];
	        this.privacy = source["privacy"];
	        this.profileIconId = source["profileIconId"];
	        this.puuid = source["puuid"];
	        this.rerollPoints = this.convertValues(source["rerollPoints"], Object);
	        this.summonerId = source["summonerId"];
	        this.summonerLevel = source["summonerLevel"];
	        this.tagLine = source["tagLine"];
	        this.unnamed = source["unnamed"];
	        this.xpSinceLastLevel = source["xpSinceLastLevel"];
	        this.xpUntilNextLevel = source["xpUntilNextLevel"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

