export enum TopLevelCategory {
	Courses,
	Services,
	Books,
}

export interface ITopPageAdvantage {
	title: string;
	description: string;
	_id: string;
}

export interface Blog {
	h1: string;
	metaTitle: string;
	metaDescription: string;
	views: number;
	_id: string;
}

export interface Sravnikus {
	metaTitle: string;
	metaDescription: string;
	qas: any[];
	_id: string;
}

export interface Learningclub {
	metaTitle: string;
	metaDescription: string;
	qas: any[];
	_id: string;
}

export interface ITopPageModel {
	_id: string;
	tags: string[];
	secondCategory: string;
	alias: string;
	title: string;
	category: string;
	seoText?: string;
	tagsTitle: string;
	metaTitle: string;
	metaDescription: string;
	firstCategory: TopLevelCategory;
	createdAt: Date;
	updatedAt: Date;
	__v: number;
	qas: any[];
	addresses: any[];
	categoryOn: string;
	blog: Blog;
	sravnikus: Sravnikus;
	learningclub: Learningclub;
}