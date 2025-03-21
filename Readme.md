# ReBAC kata

This is a kata to learn about Relation-Based Access Control (ReBAC) (see: https://en.wikipedia.org/wiki/Relationship-based_access_control).

## Subject

You are a classified ad company.
The user ask for a feature allowing to give someone a permission to edit one of their own ads.

## Follow-up

Bonus requirements:

* Add a notion of an organization. Only the members of the organization are allowed to delegate
to other members of the organization.
* Add a super admin. The super admin can create other admins. Admins can give delegations to anyone across organisations.
* Add an option to make the delegated authorization temporary.

## Constraint

The permission logic code should not leak into the Ad code and the User management code except for checking permissions and declaring relations.