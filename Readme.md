# Relation Based Access Control (ReBAC) kata

This is a kata to learn about Relation-Based Access Control (ReBAC) (see: https://en.wikipedia.org/wiki/Relationship-based_access_control).

## Subject

You are a developer in a classified ad company.
Your task is to implement a feature to allow users to give the permission
to other users to edit their ads.

## Hints

You can activate the proposed automated tests in order to gradually implement
the solution.

## Follow-up

Bonus requirements:

* Add a notion of an organization. The members of the organization are only allowed to delegate to other members of the organization.
* Add a super admin. The super admin can create other admins. Admins can give delegations to anyone across organisations.
* Add an option to make the delegated authorization temporary.

## Constraint

The permission logic code should not leak into the Ad code and the User management code except for checking permissions and declaring relations.

## Notes

The User domain and Ad domain should remain decoupled as much as possible.
ReBAC is a wonderful tool for having loosely coupled domains.
To convince yourself about this, you can contrast a ReBAC solution with a solution without ReBAC.
